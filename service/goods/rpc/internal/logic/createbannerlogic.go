package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBannerLogic {
	return &CreateBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBannerLogic) CreateBanner(in *pb.BannerRequest) (*pb.BannerResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.BannerResponse{}, nil
}
