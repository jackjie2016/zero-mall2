package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubCategoryLogic {
	return &GetSubCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取子分类
func (l *GetSubCategoryLogic) GetSubCategory(in *pb.CategoryListRequest) (*pb.SubCategoryListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SubCategoryListResponse{}, nil
}
