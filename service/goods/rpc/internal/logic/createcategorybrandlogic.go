package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	model "zero-mal/service/goods/model/gorm"

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

	//var category model.Category
	//if result := global.DB.First(&category, in.CategoryId); result.RowsAffected == 0 {
	//	return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	//}
	//
	//var brand model.Brands
	//if result := global.DB.First(&brand, in.BrandId); result.RowsAffected == 0 {
	//	return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	//}

	if _, err := l.svcCtx.BrandsModel.FindOne(l.ctx, int64(in.BrandId)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	if _, err := l.svcCtx.CategoryModel.FindOne(l.ctx, int64(in.CategoryId)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "分类不存在")
	}

	categoryBrand := model.GoodsCategoryBrand{
		CategoryId: in.CategoryId,
		BrandId:    in.BrandId,
	}

	//global.DB.Save(&categoryBrand)
	if err := l.svcCtx.GoodsCategoryBrandModel.Insert(l.ctx, &categoryBrand); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "绑定失败")
	}
	return &pb.CategoryBrandResponse{Id: categoryBrand.Id}, nil
}
