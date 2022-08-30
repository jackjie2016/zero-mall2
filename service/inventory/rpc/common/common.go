package common

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"zero-mal/global"
	model "zero-mal/service/inventory/model/gorm"
)

func AutoReback(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	type OrderInfo struct {
		OrderSn string
	}
	for i := range msgs {
		//	既然是归还库存，那么我应该具体的知道每件商品应该归还多少，但是有一个问题是说明？重复归还的问题
		//	所以说这个接口应该确保幂等性，你不能因为消息的重复发送导致应该订单的库存多少归还，没有扣减的库存你别归还
		//	如果确保这些都没有问题，新建一张表，这张表记录了详细的订单扣减细节，以及归还细节
		var orderInfo OrderInfo
		err := json.Unmarshal(msgs[i].Body, &orderInfo)
		if err != nil {
			zap.S().Errorf("json 解析失败：%v", err.Error())
			return consumer.ConsumeSuccess, nil
		}

		//库存归还，更改记录表
		tx := global.DB.Begin()
		var sellerDetail model.StockSellDetail
		if result := tx.Where(&model.StockSellDetail{OrderSn: orderInfo.OrderSn, Status: 1}).First(&sellerDetail); result.RowsAffected == 0 {
			return consumer.ConsumeSuccess, nil
		}
		for _, orderGoods := range sellerDetail.Detail {
			//更新
			if result := tx.Model(&model.Inventory{}).Where(&model.Inventory{GoodsID: orderGoods.GoodsID}).Update("stocks", gorm.Expr("stocks+?", orderGoods.Nums)); result.RowsAffected == 0 {
				tx.Rollback()
				return consumer.ConsumeRetryLater, nil //重试
			}
		}

		sellerDetail.Status = 2
		//更新sellerstockdetail 记录
		if result := tx.Model(&model.StockSellDetail{}).Where(&model.StockSellDetail{OrderSn: orderInfo.OrderSn, Status: 1}).Update("status", 2); result.RowsAffected == 0 {
			tx.Rollback()
			return consumer.ConsumeRetryLater, nil //重试
		}
		tx.Commit()
	}
	return consumer.ConsumeSuccess, nil
}
