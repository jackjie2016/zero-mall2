package logic

import (
	"context"
	"encoding/json"
	"zero-mal/global"
	model "zero-mal/service/goods/model/gorm"

	"zero-mal/service/goods/rpc/goods_pb"
	"zero-mal/service/goods/rpc/internal/svc"

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
func (l *GetAllCategorysListLogic) GetAllCategorysList(in *goods_pb.Empty) (*goods_pb.CategoryListResponse, error) {
	// todo: add your logic here and delete this line
	var categorys []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(&categorys)
	return &goods_pb.CategoryListResponse{JsonData: string(b)}, nil

}
