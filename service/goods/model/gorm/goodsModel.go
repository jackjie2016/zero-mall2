package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ GoodsModel = (*customGoodsModel)(nil)

type (
	// GoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsModel.
	GoodsModel interface {
		goodsModel
	}

	customGoodsModel struct {
		*defaultGoodsModel
	}
)

// NewGoodsModel returns a model for the database table.
func NewGoodsModel(conn *gorm.DB, c cache.CacheConf) GoodsModel {
	return &customGoodsModel{
		defaultGoodsModel: newGoodsModel(conn, c),
	}
}
