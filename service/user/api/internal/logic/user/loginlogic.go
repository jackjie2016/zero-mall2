package user

import (
	"context"
	"fmt"
	"zero-mal/common/tool"

	//"github.com/dgrijalva/jwt-go"
	errorx "zero-mal/common/error"
	"zero-mal/service/user/api/internal/svc"
	"zero-mal/service/user/api/internal/types"
	pb "zero-mal/service/user/rpc/user_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.Loginrequest) (resp *types.Loginresponse, err error) {

	//1 查找用户是否存在
	var response *pb.UserInfoResponse
	//图像验证码验证，同一个包下面的变量可以共用
	var store = tool.RedisStore{}
	if !store.Verify(req.CaptchaId, req.Captcha, true) {
		return nil, errorx.NewDefaultError("验证码错误")
	}

	if response, err = l.svcCtx.UserRpc.GetUserByMobile(l.ctx, &pb.MobileRequest{Mobile: req.Mobile}); err != nil {
		return nil, errorx.NewDefaultError("用户不存在")
	}

	//2、判断密码
	fmt.Println(req.Password, response.Password)

	var Check *pb.CheckResponse
	if Check, err = l.svcCtx.UserRpc.CheckPassWord(l.ctx, &pb.CheckInfo{
		Password:          req.Password,
		EncryptedPassword: response.Password,
	}); err != nil {
		return nil, errorx.NewDefaultError("校验失败")
	}

	//m := make(map[string]string)
	//m["22"] = "333"
	//m["225"] = "3335"
	//l.svcCtx.Cache.Set("key2", m)

	if !Check.Success {
		return nil, errorx.NewDefaultError("密码错误")
	}
	//3、生成token
	var GenerateTokenResp *pb.GenerateTokenResp
	if GenerateTokenResp, err = l.svcCtx.UserRpc.GenerateToken(l.ctx, &pb.GenerateTokenReq{UserId: int64(response.Id), Role: 1}); err != nil {
		return nil, errorx.NewDefaultError("用户不存在")
	}

	// todo: add your logic here and delete this line
	return &types.Loginresponse{
		AccessToken:  GenerateTokenResp.AccessToken,
		AccessExpire: GenerateTokenResp.AccessExpire,
		RefreshAfter: GenerateTokenResp.RefreshAfter,
	}, nil
}

//func GenToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64, userId uint32) (string, error) {
//	//
//	//claims := make(jwt.MapClaims)
//	//claims["exp"] = iat + seconds
//	//claims["iat"] = iat
//	//claims["userId"] = userId
//	//for k, v := range payloads {
//	//	claims[k] = v
//	//}
//	//token := jwt.New(jwt.SigningMethodHS256)
//	//token.Claims = claims
//	//return token.SignedString([]byte(secretKey))
//}
