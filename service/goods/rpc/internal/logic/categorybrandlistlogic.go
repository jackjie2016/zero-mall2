package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryBrandListLogic {
	return &CategoryBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌分类
func (l *CategoryBrandListLogic) CategoryBrandList(in *pb.CategoryBrandFilterRequest) (*pb.CategoryBrandListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CategoryBrandListResponse{}, nil
}
