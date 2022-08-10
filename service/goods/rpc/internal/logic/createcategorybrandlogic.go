package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryBrandLogic {
	return &CreateCategoryBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryBrandLogic) CreateCategoryBrand(in *pb.CategoryBrandRequest) (*pb.CategoryBrandResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.CategoryBrandResponse{}, nil
}
