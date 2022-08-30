package logic

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/common/tool"
	"zero-mal/common/xerr"
	"zero-mal/global"
	Grommodel "zero-mal/service/user/model/gorm"
	"zero-mal/service/user/rpc/internal/svc"
	pb "zero-mal/service/user/rpc/user_pb"
	"zero-mal/service/user/rpc/usercenter"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrGenerateTokenError = xerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = xerr.NewErrMsg("账号或密码不正确")

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *pb.CreateUserInfo) (*pb.GenerateTokenResp, error) {
	// todo: add your logic here and delete this line
	//1、判断根据手机号用户是否存在
	user, err := l.svcCtx.UserGormModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != Grommodel.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user  fail"), "err : %v , in : %+v", err, in)
	}

	if user != nil {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
		//return nil, status.Errorf(codes.Internal, "创建用户失败")
		//return nil, errors.Wrapf(ErrUserNoExistsError, "Mobile:%d", in.Mobile)
	}

	//2 、密码加密
	options := &password.Options{6, 100, 30, sha512.New}
	salt, encodedPwd := password.Encode(in.Password, options)

	in.Password = fmt.Sprintf("$zifeng-sha512$%s$%s", salt, encodedPwd)
	if len(in.NickName) == 0 {
		in.NickName = tool.Krand(8, tool.KC_RAND_KIND_ALL)
	}

	//3、构建用户模型
	userData := &Grommodel.User{
		Mobile:   in.Mobile,
		NickName: in.NickName,
		Password: in.Password,
	}
	tx := global.DB.Begin()
	err = l.svcCtx.UserGormModel.Insert(l.ctx, userData)
	if err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "创建用户失败")
	}

	zap.S().Infof("回滚了不会执行这边了的吧")
	tx.Commit()
	//userId := userData.Id
	////2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userData.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userData.Id)
	}

	respUser := pb.GenerateTokenResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}

	return &respUser, nil
}
