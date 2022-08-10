package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌和轮播图
func (l *BrandListLogic) BrandList(in *pb.BrandFilterRequest) (*pb.BrandListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.BrandListResponse{}, nil
}
