package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ StockSellDetailModel = (*customStockSellDetailModel)(nil)

type (
	// StockSellDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStockSellDetailModel.
	StockSellDetailModel interface {
		stockSellDetailModel
	}

	customStockSellDetailModel struct {
		*defaultStockSellDetailModel
	}
)

// NewStockSellDetailModel returns a model for the database table.
func NewStockSellDetailModel(conn *gorm.DB, c cache.CacheConf) StockSellDetailModel {
	return &customStockSellDetailModel{
		defaultStockSellDetailModel: newStockSellDetailModel(conn, c),
	}
}
