package common

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"
	"zero-mal/global"
	model "zero-mal/service/order/model/gorm"
)

func OrderTimeout(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {

	for i := range msgs {
		var orderInfo model.Order

		_ = json.Unmarshal(msgs[i].Body, &orderInfo)
		if result := global.DB.Where(model.Order{OrderSn: orderInfo.OrderSn}).First(&orderInfo); result.RowsAffected == 0 {
			zap.S().Info("当前订单已处理完成")
		}
		//判断当前订单状态是否为未支付，如果是执行超时机制
		if orderInfo.Status == "PAYING" {
			tx := global.DB.Begin()

			orderInfo.Status = "TRADE_CLOSED"
			tx.Save(&orderInfo)

			p, _ := rocketmq.NewProducer(
				producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"43.242.33.9:9876"})),
				producer.WithRetry(2),
			)
			err := p.Start()
			if err != nil {
				zap.S().Errorf("start producer error: %s", err.Error())
				tx.Rollback()
				return consumer.ConsumeRetryLater, nil
			}
			msg := &primitive.Message{
				Topic: "order_reback",
				Body:  msgs[i].Body,
			}
			_, err = p.SendSync(context.Background(), msg)

			if err != nil {
				tx.Rollback()
				return consumer.ConsumeRetryLater, nil
			}

			tx.Commit()

		}
	}
	return consumer.ConsumeSuccess, nil
}
