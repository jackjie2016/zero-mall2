package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ InventoryModel = (*customInventoryModel)(nil)

type (
	// InventoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInventoryModel.
	InventoryModel interface {
		inventoryModel
	}

	customInventoryModel struct {
		*defaultInventoryModel
	}
)

// NewInventoryModel returns a model for the database table.
func NewInventoryModel(conn *gorm.DB, c cache.CacheConf) InventoryModel {
	return &customInventoryModel{
		defaultInventoryModel: newInventoryModel(conn, c),
	}
}
