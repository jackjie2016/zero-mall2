package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"zero-mal/common/xerr"
	"zero-mal/service/user/rpc/usercenter"

	"zero-mal/service/user/rpc/internal/svc"
	pb "zero-mal/service/user/rpc/user_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByMobileLogic {
	return &GetUserByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByMobileLogic) GetUserByMobile(in *pb.MobileRequest) (*pb.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.UserGormModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user  fail"), "err : %v , in : %+v", err, in)
	}

	if user == nil {
		//return nil, status.Errorf(codes.Internal, "创建用户失败")
		return nil, errors.Wrapf(ErrUserNoExistsError, "Mobile:%s", in.Mobile)
	}
	var respUser usercenter.UserInfoResponse

	//不同类型的结构体不能直接赋值
	_ = copier.Copy(&respUser, user)

	if user.Birthday != nil {
		respUser.BirthDay = uint64(user.Birthday.Unix())
	}

	return &respUser, nil

}
