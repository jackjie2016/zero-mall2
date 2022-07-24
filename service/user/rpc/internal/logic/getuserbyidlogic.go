package logic

import (
	"context"

	"zero-mal/service/user/rpc/internal/svc"
	"zero-mal/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *pb.GetUserByIdReq) (*pb.GetUserByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserByIdResp{
		User: &pb.User{
			Id:       in.Id,
			Mobile:   "15958615799",
			Password: "",
			NickName: "",
			HeadUrl:  "",
			Birthday: 0,
			Address:  "",
			Desc:     "",
			Gender:   "",
			Role:     0,
		},
	}, nil
}
