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

	brandListResponse := pb.BrandListResponse{}

	//var category *model.Category
	var err error
	//if result := global.DB.Find(&category, in.Id).First(&category); result.RowsAffected == 0 {
	//	return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	//}

	if _, err = l.svcCtx.CategoryModel.FindOne(l.ctx, int64(in.Id)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var categoryBrands []model.GoodsCategoryBrand
	//if result := global.DB.Preload("Brands").Where(&model.GoodsCategoryBrand{CategoryId: in.Id}).Find(&categoryBrands); result.RowsAffected > 0 {
	if result := global.DB.Where(&model.GoodsCategoryBrand{CategoryId: in.Id}).Find(&categoryBrands); result.RowsAffected > 0 {
		brandListResponse.Total = int32(result.RowsAffected)
	}
	var brand_temp *model.Brands
	var brandInfoResponses []*pb.BrandInfoResponse
	for _, categoryBrand := range categoryBrands {
		//查询分类，用go-zero的model
		if brand_temp, err = l.svcCtx.BrandsModel.FindOne(l.ctx, int64(categoryBrand.BrandId)); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
		}

		brandInfoResponses = append(brandInfoResponses, &pb.BrandInfoResponse{
			Id:   brand_temp.Id,
			Name: brand_temp.Name,
			Logo: brand_temp.Logo,
		})
	}

	brandListResponse.Data = brandInfoResponses

	return &brandListResponse, nil
}
