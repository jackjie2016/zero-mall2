package svc

import (
	"zero-mal/global"
	"zero-mal/service/goods/rpc/internal/config"

	model "zero-mal/service/goods/model/gorm"

	"zero-mal/service/goods/rpc/internal/initialize"
)

type ServiceContext struct {
	Config                  config.Config
	GoodsModel              model.GoodsModel
	BannerModel             model.BannerModel
	BrandsModel             model.BrandsModel
	CategoryModel           model.CategoryModel
	GoodsCategoryBrandModel model.GoodsCategoryBrandModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	//if global.DB == nil {
	//	initialize.InitDb(c.DB.DataSource)
	//}
	initialize.InitDb(c.DB.DataSource)
	initialize.InitEs(c.Es)
	return &ServiceContext{
		Config:                  c,
		GoodsModel:              model.NewGoodsModel(global.DB, c.Cache),
		BannerModel:             model.NewBannerModel(global.DB, c.Cache),
		BrandsModel:             model.NewBrandsModel(global.DB, c.Cache),
		CategoryModel:           model.NewCategoryModel(global.DB, c.Cache),
		GoodsCategoryBrandModel: model.NewGoodsCategoryBrandModel(global.DB, c.Cache),
	}
}
