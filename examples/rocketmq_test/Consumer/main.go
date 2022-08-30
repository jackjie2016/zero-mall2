package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.31.147:9876"}),
		consumer.WithGroupName("gozero"),
	)

	if err := c.Subscribe("order", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("获取到值： %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	}); err != nil {
		fmt.Printf("读取消息失败:%s", err.Error())
		panic(err.Error())
	}
	err := c.Start()
	if err != nil {
		fmt.Printf("start consuming error:%s\n", err)
		return
	}
	//不能让主goroutine退出
	time.Sleep(time.Hour)
	_ = c.Shutdown()
}
