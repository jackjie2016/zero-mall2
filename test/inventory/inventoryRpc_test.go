/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package inventory

import (
	"context"
	"fmt"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"math/rand"
	"testing"
	"time"
	proto "zero-mal/service/inventory/rpc/inventory_pb"
)

//函数中通过调用testing.T 的 Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息
var InventoryClient proto.InventoryClient
var conn *grpc.ClientConn

func init() {
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

	defer conn.Close()
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
func TestInvDetail(t *testing.T) {

	//t.Fatalf("parse messageID %s error", id)
	//
	//t.Log(msgID)
}

func TestRreback(t *testing.T) {

	//t.Fatalf("parse messageID %s error", id)
	//
	//t.Log(msgID)
}
func TestSell(t *testing.T) {

	//t.Fatalf("parse messageID %s error", id)
	//
	//t.Log(msgID)
}
func TestSetinv(t *testing.T) {
	tests := []struct {
		name      string
		GoodsInfo []*proto.GoodsInvInfo
		OrderSn   string
		err       error
	}{
		{
			name: "ok",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(rand.Int31n(90) * 2),
			err:     nil,
		},
		{
			name: "Inventory_Internal",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(rand.Int31n(90) * 2),
			err:     status.Errorf(codes.Internal, "保存库存扣减历史失败"),
		},
		{
			name: "order_sn_inserted",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(rand.Int31n(90) * 2),
			err:     status.Errorf(codes.NotFound, "没有库存信息"),
		},
		{
			name: "goods_NotFound",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(rand.Int31n(90) * 3),
			err:     status.Errorf(codes.NotFound, "没有库存信息"),
		},
		{
			name: "ResourceExhausted",
			GoodsInfo: []*proto.GoodsInvInfo{
				{GoodsId: 421, Num: 1000},
				//{GoodsId: 2, Num: 1},
			},
			OrderSn: GenerateOrderSn(rand.Int31n(90) * 4),
			err:     status.Errorf(codes.ResourceExhausted, "库存不足"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := InventoryClient.Sell(context.Background(), &proto.SellInfo{
				GoodsInfo: test.GoodsInfo,
				OrderSn:   test.OrderSn,
			})

			if err != test.err {
				t.Errorf("%v", err)
			}
		})
	}
}

// Benchmark Test
//cmd go test -bench=".*" -benchmem -count=6 .

func BenchmarkSetinv(b *testing.B) {
	for k := 0; k < b.N; k++ {

	}
}

func BenchmarkSell(b *testing.B) {
	for k := 0; k < b.N; k++ {

	}
}
