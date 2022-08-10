package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BannerListLogic {
	return &BannerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 轮播图
func (l *BannerListLogic) BannerList(in *pb.Empty) (*pb.BannerListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.BannerListResponse{}, nil
}
