package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AddressModel = (*customAddressModel)(nil)

type (
	// AddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAddressModel.
	AddressModel interface {
		addressModel
	}

	customAddressModel struct {
		*defaultAddressModel
	}
)

// NewAddressModel returns a model for the database table.
func NewAddressModel(conn sqlx.SqlConn, c cache.CacheConf) AddressModel {
	return &customAddressModel{
		defaultAddressModel: newAddressModel(conn, c),
	}
}
