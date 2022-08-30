package main

import (
	"context"
	"fmt"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
	"sync"
	"time"
	"zero-mal/common/tool"
	proto "zero-mal/service/inventory/rpc/inventory_pb"
)

var InventoryClient proto.InventoryClient
var conn *grpc.ClientConn

func TestInvDetail(GoodsId int32) {
	res, err := InventoryClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: GoodsId,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func TestSetInv(GoodsId, Num int32) {
	_, err := InventoryClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: GoodsId,
		Num:     Num,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("设置成功")
}

func TestSell(wg *sync.WaitGroup, i int32) {
	worker, _ := tool.NewOrderWorker(0)
	defer wg.Done()
	_, err := InventoryClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{GoodsId: 421,

				Num: 1},
			//{GoodsId: 2, Num: 1},
		},
		OrderSn: fmt.Sprintf("%d", worker.NextOrderId()),
	})
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()
	if err != nil {
		panic(err)
	}

	fmt.Println(i, "扣减成功")
}
func Init() {
	//ip, _ := utils.ExternalIP()

	//IP := ip.String()
	//fmt.Printf("%s:8002", IP)

	//********方式二;在此处配置  ******
	//********重试的话需要做好幂等性******
	retryOpt := []retry.CallOption{
		retry.WithMax(2), //重试3次
		retry.WithPerRetryTimeout(1 * time.Second),                                // 超过1s就要重试
		retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Unavailable), // 哪些状态码重试
	}

	var err error
	conn, err = grpc.Dial("127.0.0.1:8002",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(retryOpt...)),
	)
	if err != nil {
		panic(err)
	}
	InventoryClient = proto.NewInventoryClient(conn)
}
func GenerateOrderSn(userId int32) string {
	//订单号的生成规则
	/*
		年月日时分秒+用户id+2位随机数
	*/
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(),
		userId, rand.Intn(90)+10,
	)
	return orderSn
}
func main() {

	Init()
	//TestCreateUser()
	//TestSetInv(2,2)

	//TestGetGoodsDetail()
	var wg sync.WaitGroup
	wg.Add(999)
	for i := 0; i < 999; i++ {
		go func(k int) {
			TestSell(&wg, int32(k))
		}(i)
	}
	wg.Wait()

	conn.Close()
}
