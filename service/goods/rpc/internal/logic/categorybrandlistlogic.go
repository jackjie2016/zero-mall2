package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/common/tool"
	"zero-mal/global"
	model "zero-mal/service/goods/model/gorm"

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
	var categoryBrands []model.GoodsCategoryBrand
	categoryBrandListResponse := pb.CategoryBrandListResponse{}

	var total int64
	global.DB.Model(&model.GoodsCategoryBrand{}).Count(&total)
	categoryBrandListResponse.Total = int32(total)

	global.DB.Scopes(tool.Paginate(int(in.Pages), int(in.PagePerNums))).Find(&categoryBrands)

	//Preload("Category").Preload("Brands").
	var brand_temp *model.Brands
	var categroy_temp *model.Category
	var err error
	var categoryResponses []*pb.CategoryBrandResponse
	for _, categoryBrand := range categoryBrands {
		//查询分类，用go-zero的model
		if brand_temp, err = l.svcCtx.BrandsModel.FindOne(l.ctx, int64(categoryBrand.BrandId)); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
		}

		if categroy_temp, err = l.svcCtx.CategoryModel.FindOne(l.ctx, int64(categoryBrand.CategoryId)); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "分类不存在")
		}

		categoryResponses = append(categoryResponses, &pb.CategoryBrandResponse{
			Id: categoryBrand.Id,
			Category: &pb.CategoryInfoResponse{
				Id:             categroy_temp.Id,
				Name:           categroy_temp.Name,
				Level:          categroy_temp.Level,
				IsTab:          categroy_temp.IsTab,
				ParentCategory: categroy_temp.ParentCategoryID,
			},
			Brand: &pb.BrandInfoResponse{
				Id:   brand_temp.Id,
				Name: brand_temp.Name,
				Logo: brand_temp.Logo,
			},
		})
	}

	categoryBrandListResponse.Data = categoryResponses
	return &categoryBrandListResponse, nil

}
