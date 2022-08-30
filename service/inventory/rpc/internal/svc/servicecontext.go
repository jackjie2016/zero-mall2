package svc

import (
	"zero-mal/global"
	model "zero-mal/service/inventory/model/gorm"
	"zero-mal/service/inventory/rpc/internal/config"
	"zero-mal/service/inventory/rpc/internal/initialize"
)

type ServiceContext struct {
	Config               config.Config
	InventoryModel       model.InventoryModel
	StockSellDetailModel model.StockSellDetailModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	initialize.InitDb(c.DB.DataSource)

	return &ServiceContext{
		Config:               c,
		InventoryModel:       model.NewInventoryModel(global.DB, c.Cache),
		StockSellDetailModel: model.NewStockSellDetailModel(global.DB, c.Cache),
	}
}
