// Code generated by goctl. DO NOT EDIT!

package gorm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stringx"
	"gorm.io/gorm"
)

var (
	deliveryFieldNames          = builder.RawFieldNames(&Delivery{})
	deliveryRows                = strings.Join(deliveryFieldNames, ",")
	deliveryRowsExpectAutoSet   = strings.Join(stringx.Remove(deliveryFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	deliveryRowsWithPlaceHolder = strings.Join(stringx.Remove(deliveryFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheGoZeroMallDeliveryIdPrefix = "cache:goZeroMall:delivery:id:"
)

type (
	deliveryModel interface {
		Insert(ctx context.Context, data *Delivery) error
		FindOne(ctx context.Context, id int64) (*Delivery, error)
		Update(ctx context.Context, data *Delivery) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDeliveryModel struct {
		gormc.CachedConn
		table string
	}

	Delivery struct {
		BaseModel

		GoodsID sql.NullInt64  `gorm:"column:goods_id"`
		Nums    sql.NullInt64  `gorm:"column:nums"`
		OrderSn sql.NullString `gorm:"column:order_sn"`
		Status  sql.NullInt64  `gorm:"column:status"`
	}
)

func newDeliveryModel(conn *gorm.DB, c cache.CacheConf) *defaultDeliveryModel {
	return &defaultDeliveryModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`delivery`",
	}
}

func (m *defaultDeliveryModel) Insert(ctx context.Context, data *Delivery) error {
	goZeroMallDeliveryIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallDeliveryIdPrefix, data.Id)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Save(&data).Error
	}, goZeroMallDeliveryIdKey)
	return err
}

func (m *defaultDeliveryModel) FindOne(ctx context.Context, id int64) (*Delivery, error) {
	goZeroMallDeliveryIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallDeliveryIdPrefix, id)
	var resp Delivery
	err := m.QueryCtx(ctx, &resp, goZeroMallDeliveryIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Delivery{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDeliveryModel) Update(ctx context.Context, data *Delivery) error {
	goZeroMallDeliveryIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallDeliveryIdPrefix, data.Id)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Save(data).Error
	}, goZeroMallDeliveryIdKey)
	return err
}

func (m *defaultDeliveryModel) Delete(ctx context.Context, id int64) error {
	goZeroMallDeliveryIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallDeliveryIdPrefix, id)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Delete(&Delivery{}, id).Error
	}, goZeroMallDeliveryIdKey)
	return err
}

func (m *defaultDeliveryModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGoZeroMallDeliveryIdPrefix, primary)
}

func (m *defaultDeliveryModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&Delivery{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultDeliveryModel) tableName() string {
	return m.table
}
