package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllCategorysListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllCategorysListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllCategorysListLogic {
	return &GetAllCategorysListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品分类
func (l *GetAllCategorysListLogic) GetAllCategorysList(in *pb.Empty) (*pb.CategoryListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CategoryListResponse{}, nil
}
