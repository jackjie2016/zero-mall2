// Code generated by goctl. DO NOT EDIT!

package gorm

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stringx"
	"gorm.io/gorm"
)

var (
	stockSellDetailFieldNames          = builder.RawFieldNames(&StockSellDetail{})
	stockSellDetailRows                = strings.Join(stockSellDetailFieldNames, ",")
	stockSellDetailRowsExpectAutoSet   = strings.Join(stringx.Remove(stockSellDetailFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	stockSellDetailRowsWithPlaceHolder = strings.Join(stringx.Remove(stockSellDetailFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheGoZeroMallStockSellDetailIdPrefix      = "cache:goZeroMall:stockSellDetail:id:"
	cacheGoZeroMallStockSellDetailOrderSnPrefix = "cache:goZeroMall:stockSellDetail:orderSn:"
)

type (
	stockSellDetailModel interface {
		Insert(ctx context.Context, data *StockSellDetail) error
		FindOne(ctx context.Context, id int64) (*StockSellDetail, error)
		FindOneByOrderSn(ctx context.Context, orderSn sql.NullString) (*StockSellDetail, error)
		Update(ctx context.Context, data *StockSellDetail) error
		Delete(ctx context.Context, id int64) error
	}

	defaultStockSellDetailModel struct {
		gormc.CachedConn
		table string
	}

	StockSellDetail struct {
		BaseModel

		OrderSn string `gorm:"column:order_sn"`
		Status  int32  `gorm:"column:status"`

		Detail GoodsDetailList `gorm:"type:varchar(200)"` //商品详情
	}

	GoodsDetail struct {
		GoodsID int32
		Nums    int32
	}

	GoodsDetailList []GoodsDetail
)

func (g GoodsDetailList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GoodsDetailList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

func newStockSellDetailModel(conn *gorm.DB, c cache.CacheConf) *defaultStockSellDetailModel {
	return &defaultStockSellDetailModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`stock_sell_detail`",
	}
}

func (m *defaultStockSellDetailModel) Insert(ctx context.Context, data *StockSellDetail) error {
	goZeroMallStockSellDetailIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallStockSellDetailIdPrefix, data.Id)
	goZeroMallStockSellDetailOrderSnKey := fmt.Sprintf("%s%v", cacheGoZeroMallStockSellDetailOrderSnPrefix, data.OrderSn)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Save(&data).Error
	}, goZeroMallStockSellDetailIdKey, goZeroMallStockSellDetailOrderSnKey)
	return err
}

func (m *defaultStockSellDetailModel) FindOne(ctx context.Context, id int64) (*StockSellDetail, error) {
	goZeroMallStockSellDetailIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallStockSellDetailIdPrefix, id)
	var resp StockSellDetail
	err := m.QueryCtx(ctx, &resp, goZeroMallStockSellDetailIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&StockSellDetail{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultStockSellDetailModel) FindOneByOrderSn(ctx context.Context, orderSn sql.NullString) (*StockSellDetail, error) {
	goZeroMallStockSellDetailOrderSnKey := fmt.Sprintf("%s%v", cacheGoZeroMallStockSellDetailOrderSnPrefix, orderSn)
	var resp StockSellDetail
	err := m.QueryRowIndexCtx(ctx, &resp, goZeroMallStockSellDetailOrderSnKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&StockSellDetail{}).Where("`order_sn` = ?", orderSn).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStockSellDetailModel) Update(ctx context.Context, data *StockSellDetail) error {
	goZeroMallStockSellDetailIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallStockSellDetailIdPrefix, data.Id)
	goZeroMallStockSellDetailOrderSnKey := fmt.Sprintf("%s%v", cacheGoZeroMallStockSellDetailOrderSnPrefix, data.OrderSn)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Save(data).Error
	}, goZeroMallStockSellDetailIdKey, goZeroMallStockSellDetailOrderSnKey)
	return err
}

func (m *defaultStockSellDetailModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	goZeroMallStockSellDetailIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallStockSellDetailIdPrefix, id)
	goZeroMallStockSellDetailOrderSnKey := fmt.Sprintf("%s%v", cacheGoZeroMallStockSellDetailOrderSnPrefix, data.OrderSn)
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Delete(&StockSellDetail{}, id).Error
	}, goZeroMallStockSellDetailIdKey, goZeroMallStockSellDetailOrderSnKey)
	return err
}

func (m *defaultStockSellDetailModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGoZeroMallStockSellDetailIdPrefix, primary)
}

func (m *defaultStockSellDetailModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&StockSellDetail{}).Where("`id` = ?", primary).Take(v).Error
}
