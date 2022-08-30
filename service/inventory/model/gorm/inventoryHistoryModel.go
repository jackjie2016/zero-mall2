package gorm

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ InventoryHistoryModel = (*customInventoryHistoryModel)(nil)

type (
	// InventoryHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInventoryHistoryModel.
	InventoryHistoryModel interface {
		inventoryHistoryModel
	}

	customInventoryHistoryModel struct {
		*defaultInventoryHistoryModel
	}
)

// NewInventoryHistoryModel returns a model for the database table.
func NewInventoryHistoryModel(conn *gorm.DB, c cache.CacheConf) InventoryHistoryModel {
	return &customInventoryHistoryModel{
		defaultInventoryHistoryModel: newInventoryHistoryModel(conn, c),
	}
}
