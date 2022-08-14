package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ BrandsModel = (*customBrandsModel)(nil)

type (
	// BrandsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBrandsModel.
	BrandsModel interface {
		brandsModel
	}

	customBrandsModel struct {
		*defaultBrandsModel
	}
)

// NewBrandsModel returns a model for the database table.
func NewBrandsModel(conn *gorm.DB, c cache.CacheConf) BrandsModel {
	return &customBrandsModel{
		defaultBrandsModel: newBrandsModel(conn, c),
	}
}
