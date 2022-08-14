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
	categoryListResponse := pb.SubCategoryListResponse{}

	var category model.Category
	if result := global.DB.First(&category, in.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	categoryListResponse.Info = &pb.CategoryInfoResponse{
		Id:             category.Id,
		Name:           category.Name,
		Level:          int32(category.Level),
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategorys []model.Category
	var subCategoryResponse []*pb.CategoryInfoResponse
	//preloads := "SubCategory"
	//if category.Level == 1 {
	//	preloads = "SubCategory.SubCategory"
	//}
	global.DB.Where(&model.Category{ParentCategoryID: in.Id}).Find(&subCategorys)

	for _, subCategory := range subCategorys {
		subCategoryResponse = append(subCategoryResponse, &pb.CategoryInfoResponse{
			Id:             subCategory.Id,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}

	categoryListResponse.SubCategorys = subCategoryResponse
	return &categoryListResponse, nil

}
