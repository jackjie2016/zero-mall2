package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"

	"gorm.io/gorm"
)

var _ GoodsCategoryBrandModel = (*customGoodsCategoryBrandModel)(nil)

type (
	// GoodsCategoryBrandModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsCategoryBrandModel.
	GoodsCategoryBrandModel interface {
		goodsCategoryBrandModel
	}

	customGoodsCategoryBrandModel struct {
		*defaultGoodsCategoryBrandModel
	}
)

// NewGoodsCategoryBrandModel returns a model for the database table.
func NewGoodsCategoryBrandModel(conn *gorm.DB, c cache.CacheConf) GoodsCategoryBrandModel {
	return &customGoodsCategoryBrandModel{
		defaultGoodsCategoryBrandModel: newGoodsCategoryBrandModel(conn, c),
	}
}
