package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CartModel = (*customCartModel)(nil)

type (
	// CartModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCartModel.
	CartModel interface {
		cartModel
	}

	customCartModel struct {
		*defaultCartModel
	}
)

// NewCartModel returns a model for the database table.
func NewCartModel(conn *gorm.DB, c cache.CacheConf) CartModel {
	return &customCartModel{
		defaultCartModel: newCartModel(conn, c),
	}
}
