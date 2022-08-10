package logic

import (
	"context"
	"github.com/pkg/errors"
	"zero-mal/common/xerr"
	model "zero-mal/service/user/model/genModel"
	"zero-mal/service/user/rpc/internal/svc"
	"zero-mal/service/user/rpc/pb"
	"zero-mal/service/user/rpc/usercenter"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//func ModelToResponse(user model.User) pb.GetUserByIdResp {
//	//UserInfoRsp := pb.GetUserByIdResp{
//	//	Id:       uint32(user.ID),
//	//	Password: user.Password,
//	//	Mobile:   user.Mobile,
//	//	NickName: user.NickName,
//	//	Gender:   user.Gender,
//	//	Role:     int32(user.Role),
//	//}
//	//UserInfoRsp := pb.GetUserByIdResp{User: }
//	//if user.Birthday != nil {
//	//	UserInfoRsp.BirthDay = uint64(user.Birthday.Unix())
//	//}
//
//	return &pb.GetUserByIdResp{}
//}
func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user  fail"), "err : %v , in : %+v", err, in)
	}

	if user == nil {
		//return nil, status.Errorf(codes.Internal, "创建用户失败")
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.Id)
	}
	var respUser usercenter.User
	_ = copier.Copy(&respUser, user)

	if user.Birthday != nil {
		respUser.Birthday = uint64(user.Birthday.Unix())
	}

	return &usercenter.GetUserByIdResp{
		User: &respUser,
	}, nil

}
