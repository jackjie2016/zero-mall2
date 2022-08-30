package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"

	"gorm.io/gorm"
)

var _ DeliveryModel = (*customDeliveryModel)(nil)

type (
	// DeliveryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeliveryModel.
	DeliveryModel interface {
		deliveryModel
	}

	customDeliveryModel struct {
		*defaultDeliveryModel
	}
)

// NewDeliveryModel returns a model for the database table.
func NewDeliveryModel(conn *gorm.DB, c cache.CacheConf) DeliveryModel {
	return &customDeliveryModel{
		defaultDeliveryModel: newDeliveryModel(conn, c),
	}
}
