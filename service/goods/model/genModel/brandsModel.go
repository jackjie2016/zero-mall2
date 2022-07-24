package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
func NewBrandsModel(conn sqlx.SqlConn, c cache.CacheConf) BrandsModel {
	return &customBrandsModel{
		defaultBrandsModel: newBrandsModel(conn, c),
	}
}
