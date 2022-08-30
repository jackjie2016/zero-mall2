package main

import (
	"flag"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"go.uber.org/zap"
	"zero-mal/service/order/rpc/common"

	"zero-mal/service/order/rpc/internal/config"
	"zero-mal/service/order/rpc/internal/server"
	"zero-mal/service/order/rpc/internal/svc"
	"zero-mal/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "service/order/rpc/etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewOrderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	//rocketMq 消费消息
	mq_client, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("zero-order"), //分布式用到，同一种实例
		consumer.WithNsResolver(primitive.NewPassthroughResolver(c.Rocketmq.Hosts)),
	)
	err := mq_client.Subscribe("order_timeout", consumer.MessageSelector{}, common.OrderTimeout)
	if err != nil {
		zap.S().Errorf("生成 consumer 失败: %s", err.Error())
		return
	}
	// Note: start after subscribe
	err = mq_client.Start()
	if err != nil {
		zap.S().Errorf("rocketmq 启动失败")
		return
	}
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
