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

package goods

import (
	"context"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"testing"
	"time"
	proto "zero-mal/service/goods/rpc/goods_pb"
)

//函数中通过调用testing.T 的 Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息
var GoodsClient proto.GoodsClient
var conn *grpc.ClientConn

func init() {

	retryOpt := []retry.CallOption{
		retry.WithMax(2), //重试3次
		retry.WithPerRetryTimeout(1 * time.Second),                                // 超过1s就要重试
		retry.WithCodes(codes.Unknown, codes.DeadlineExceeded, codes.Unavailable), // 哪些状态码重试
	}

	var err error
	conn, err = grpc.Dial("127.0.0.1:8001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(retryOpt...)),
	)
	if err != nil {
		panic(err)
	}
	GoodsClient = proto.NewGoodsClient(conn)

	defer conn.Close()
}

func TestCreateGoods(t *testing.T) {
	tests := []struct {
		name      string
		GoodsInfo *proto.CreateGoodsInfo
		err       error
	}{
		{
			name:      "ok",
			GoodsInfo: &proto.CreateGoodsInfo{},
			err:       nil,
		},
		{
			name:      "creater_Internal",
			GoodsInfo: &proto.CreateGoodsInfo{},
			err:       status.Errorf(codes.Internal, "保存失败"),
		},
		{
			name:      "repeated_inserted",
			GoodsInfo: &proto.CreateGoodsInfo{},
			err:       status.Errorf(codes.NotFound, "重复插入"),
		},
		{
			name:      "cate_NotFound",
			GoodsInfo: &proto.CreateGoodsInfo{},
			err:       status.Errorf(codes.NotFound, "分类不存在插入"),
		},
		{
			name:      "brand_notfound",
			GoodsInfo: &proto.CreateGoodsInfo{},
			err:       status.Errorf(codes.NotFound, "品牌不存在插入"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := GoodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{})

			if err != test.err {
				t.Errorf("%v", err)
			}
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
