package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryBrandListLogic {
	return &GetCategoryBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过category获取brands
func (l *GetCategoryBrandListLogic) GetCategoryBrandList(in *pb.CategoryInfoRequest) (*pb.BrandListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.BrandListResponse{}, nil
}
