package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ BannerModel = (*customBannerModel)(nil)

type (
	// BannerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBannerModel.
	BannerModel interface {
		bannerModel
	}

	customBannerModel struct {
		*defaultBannerModel
	}
)

// NewBannerModel returns a model for the database table.
func NewBannerModel(conn *gorm.DB, c cache.CacheConf) BannerModel {
	return &customBannerModel{
		defaultBannerModel: newBannerModel(conn, c),
	}
}
