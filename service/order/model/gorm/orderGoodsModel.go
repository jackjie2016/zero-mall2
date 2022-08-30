package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"

	"gorm.io/gorm"
)

var _ OrderGoodsModel = (*customOrderGoodsModel)(nil)

type (
	// OrderGoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderGoodsModel.
	OrderGoodsModel interface {
		orderGoodsModel
	}

	customOrderGoodsModel struct {
		*defaultOrderGoodsModel
	}
)

// NewOrderGoodsModel returns a model for the database table.
func NewOrderGoodsModel(conn *gorm.DB, c cache.CacheConf) OrderGoodsModel {
	return &customOrderGoodsModel{
		defaultOrderGoodsModel: newOrderGoodsModel(conn, c),
	}
}
