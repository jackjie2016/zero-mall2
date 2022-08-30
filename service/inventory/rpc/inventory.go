package main

import (
	"context"
	"flag"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"zero-mal/service/inventory/rpc/common"
	"zero-mal/service/inventory/rpc/internal/initialize"

	"zero-mal/service/inventory/rpc/internal/config"
	"zero-mal/service/inventory/rpc/internal/server"
	"zero-mal/service/inventory/rpc/internal/svc"
	pb "zero-mal/service/inventory/rpc/inventory_pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "service/inventory/rpc/etc/inventory-dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewInventoryServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterInventoryServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	initialize.InitSentinel()
	Interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		//限流
		ee, eb := sentinel.Entry("limit_WarmUp", sentinel.WithTrafficType(base.Inbound))
		if eb != nil {
			fmt.Println("限流了")
			return nil, status.Errorf(codes.Internal, "限流了")
		} else {
			fmt.Println("正常")
			ee.Exit()
		}
		fmt.Println("接收一个请求")
		res, err := handler(ctx, req)
		fmt.Println("请求结束")
		return res, err
	}
	s.AddUnaryInterceptors(Interceptor)
	defer s.Stop()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	startMq := make(chan struct{})
	go func() {
		<-startMq
		time.Sleep(time.Second * 5)
		//rocketMq 消费消息
		mq_client, _ := rocketmq.NewPushConsumer(
			consumer.WithGroupName("zero-inventory"), //分布式用到，同一种实例
			consumer.WithNsResolver(primitive.NewPassthroughResolver(c.Rocketmq.Hosts)),
		)
		err := mq_client.Subscribe("order_reback", consumer.MessageSelector{}, common.AutoReback)
		if err != nil {
			logx.Errorf("生成 consumer 失败: %s", err.Error())
			return
		}
		// Note: start after subscribe
		err = mq_client.Start()
		if err != nil {
			logx.Error("rocketmq 启动失败")
			return
		}
	}()
	startMq <- struct{}{}
	s.Start()

}
