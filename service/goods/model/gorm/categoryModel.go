package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"

	"gorm.io/gorm"
)

var _ CategoryModel = (*customCategoryModel)(nil)

type (
	// CategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoryModel.
	CategoryModel interface {
		categoryModel
	}

	customCategoryModel struct {
		*defaultCategoryModel
	}
)

// NewCategoryModel returns a model for the database table.
func NewCategoryModel(conn *gorm.DB, c cache.CacheConf) CategoryModel {
	return &customCategoryModel{
		defaultCategoryModel: newCategoryModel(conn, c),
	}
}
