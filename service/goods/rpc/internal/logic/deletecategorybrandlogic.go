package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/global"
	model "zero-mal/service/goods/model/gorm"
	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryBrandLogic {
	return &DeleteCategoryBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryBrandLogic) DeleteCategoryBrand(in *pb.CategoryBrandRequest) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	if result := global.DB.Delete(&model.GoodsCategoryBrand{}, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}

	return &pb.Empty{}, nil
}
