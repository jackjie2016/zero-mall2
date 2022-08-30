package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"zero-mal/common/ctxdata"
	"zero-mal/service/user/rpc/internal/svc"
	pb "zero-mal/service/user/rpc/user_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {

	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.UserGormModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "用户不存在")
	}

	var GenerateTokenResp *pb.GenerateTokenResp
	l.svcCtx.Cache.Get(fmt.Sprintf("login_token:%d", in.UserId), &GenerateTokenResp)

	if GenerateTokenResp != nil {
		return GenerateTokenResp, nil
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", in.UserId, err)
	}
	GenerateTokenResp = &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}
	l.svcCtx.Cache.SetWithExpire(fmt.Sprintf("login_token:%d", in.UserId), GenerateTokenResp, time.Duration(accessExpire-20))

	return GenerateTokenResp, nil

}

func getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
