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

package order

import (
	"context"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"testing"
	"time"
	proto "zero-mal/service/order/rpc/pb"
)

//函数中通过调用testing.T 的 Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息
var OrderClient proto.OrderClient
var conn *grpc.ClientConn

func init() {

	retryOpt := []retry.CallOption{
		retry.WithMax(2), //重试3次
		retry.WithPerRetryTimeout(1 * time.Second),                                // 超过1s就要重试
		retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Unavailable), // 哪些状态码重试
	}

	var err error
	conn, err = grpc.Dial("127.0.0.1:8003",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(retryOpt...)),
	)
	if err != nil {
		panic(err)
	}
	OrderClient = proto.NewOrderClient(conn)

	defer conn.Close()
}

func TestCreateOrder(t *testing.T) {
	tests := []struct {
		name         string
		OrderRequest *proto.OrderRequest
		GoodsIds     []int32
		err          error
	}{
		{
			name:         "ok",
			OrderRequest: &proto.OrderRequest{},
			err:          nil,
		},
		{
			name:         "creater_Internal",
			OrderRequest: &proto.OrderRequest{},
			GoodsIds:     []int32{2, 3},
			err:          status.Errorf(codes.Internal, "保存失败"),
		},
		{
			name:         "add_cart_Internal",
			OrderRequest: &proto.OrderRequest{},
			GoodsIds:     []int32{2, 3},
			err:          status.Errorf(codes.NotFound, "添加购物车失败"),
		},
		{
			name:         "goods_notfound",
			OrderRequest: &proto.OrderRequest{},
			GoodsIds:     []int32{2, 3},
			err:          status.Errorf(codes.NotFound, "商品没有"),
		},
		{
			name:         "repeated_inserted",
			OrderRequest: &proto.OrderRequest{},
			GoodsIds:     []int32{2, 3},
			err:          status.Errorf(codes.NotFound, "重复插入"),
		},
		{
			name:         "cart_empty",
			OrderRequest: &proto.OrderRequest{},
			GoodsIds:     []int32{2, 3},
			err:          status.Errorf(codes.NotFound, "购物车空的"),
		},
		{
			name:         "inventory_Internal",
			OrderRequest: &proto.OrderRequest{},
			GoodsIds:     []int32{2, 3},
			err:          status.Errorf(codes.NotFound, "库存扣减失败"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			//1，商品加入购物车

			for _, v := range test.GoodsIds {
				_, err := OrderClient.AddCartItem(context.Background(), &proto.CartItemRequest{GoodsId: v})
				if err != test.err {
					t.Errorf("%v", err)
				}
			}
			//2、创建订单
			_, err := OrderClient.CreateOrder(context.Background(), &proto.OrderRequest{})
			if err != test.err {
				t.Errorf("%v", err)
			}
			t.Log("pass")
		})
	}
}

// Benchmark Test
//cmd go test -bench=".*" -benchmem -count=6 .
//测试商品详情

func BenchmarkGetGoodsDetail(b *testing.B) {
	for k := 0; k < b.N; k++ {

	}
}

func BenchmarkGoodsList(b *testing.B) {
	for k := 0; k < b.N; k++ {

	}
}
