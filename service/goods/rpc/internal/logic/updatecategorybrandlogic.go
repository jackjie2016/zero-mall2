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

type UpdateCategoryBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryBrandLogic {
	return &UpdateCategoryBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryBrandLogic) UpdateCategoryBrand(in *pb.CategoryBrandRequest) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	var categoryBrand model.GoodsCategoryBrand

	if result := global.DB.First(&categoryBrand, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌分类不存在")
	}

	if _, err := l.svcCtx.BrandsModel.FindOne(l.ctx, int64(in.BrandId)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	if _, err := l.svcCtx.CategoryModel.FindOne(l.ctx, int64(in.CategoryId)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "分类不存在")
	}

	categoryBrand.CategoryId = in.CategoryId
	categoryBrand.BrandId = in.BrandId

	global.DB.Save(&categoryBrand)

	return &pb.Empty{}, nil
}
