package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/global"
)

func main() {

	//redis 分布式锁
	tx := global.DB.Begin()

	fmt.Printf("len=%d cap=%d slice=%v\n", len(GoodsDetail), cap(GoodsDetail), GoodsDetail)
	if l := len(GoodsDetail); l > 0 {
		var stockSellDetail = model.StockSellDetail{
			OrderSn: req.OrderSn,
			Status:  1,
			Detail:  GoodsDetail,
		}
		if result := tx.Create(stockSellDetail); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "保存库存扣减历史失败")
		}
	}

	zap.S().Infof("回滚了不会执行这边了的吧")
	tx.Commit()
}
