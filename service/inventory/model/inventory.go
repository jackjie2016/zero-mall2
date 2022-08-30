package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Delivery struct {
	Goods   int32 `gorm:"type:int;index"`
	Nums    int32 `gorm:"type:int"`
	OrderSn int32 `gorm:"type:varchar(200)"`
	Status  int32 `gorm:"type:int"`
}

type StockSellDetail struct {
	OrderSn string          `gorm:"type:varchar(200);index:idx_index_ordersn,unique"`
	Status  int32           `gorm:"type:int"`          //1、已扣减，2、已归还
	Detail  GoodsDetailList `gorm:"type:varchar(200)"` //商品详情
}

func (StockSellDetail) TableName() string {
	return "stock_sell_detail"
}

type GoodsDetail struct {
	Goods int32
	Nums  int32
}

type GoodsDetailList []GoodsDetail

func (g GoodsDetailList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GoodsDetailList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}
