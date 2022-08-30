package logic

import (
	"context"
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v8"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/global"
	model "zero-mal/service/inventory/model/gorm"

	"zero-mal/service/inventory/rpc/internal/svc"
	pb "zero-mal/service/inventory/rpc/inventory_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SellLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSellLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SellLogic {
	return &SellLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SellLogic) Sell(in *pb.SellInfo) (*pb.Empty, error) {
	// todo: add your logic here and delete this line

	//redis 分布式锁
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
	})
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	//redis 分布式锁
	tx := global.DB.Begin()
	var GoodsDetail model.GoodsDetailList
	//m.Lock()

	mutex := rs.NewMutex("stock_sell")
	if err := mutex.Lock(); err != nil {
		logx.Error("redis 加锁失败")
		return nil, status.Errorf(codes.Internal, "redis 加锁失败")
	}

	defer func() {
		if ok, err := mutex.Unlock(); !ok || err != nil {
			logx.Error("redis 解锁失败")

		}
	}()
	for _, goods := range in.GoodsInfo {
		//if err := mutex.Lock(); err != nil {
		//	logx.Error("redis 加锁失败")
		//	return nil, status.Errorf(codes.Internal, "redis 加锁失败")
		//}
		var inv model.Inventory
		if res := tx.Where(&model.Inventory{GoodsID: goods.GoodsId}).First(&inv); res.RowsAffected == 0 {
			logx.Infof("回滚吧")
			tx.Rollback()
			return nil, status.Errorf(codes.NotFound, "没有库存信息")
		}

		if inv.Stocks < goods.Num {
			logx.Infof("回滚吧")
			tx.Rollback()
			return nil, status.Errorf(codes.ResourceExhausted, "库存不足")
		}
		inv.Stocks -= goods.Num
		tx.Where(&model.Inventory{}).Save(&inv)
		GoodsDetail = append(GoodsDetail, model.GoodsDetail{
			GoodsID: goods.GoodsId,
			Nums:    goods.Num,
		})

		//if ok, err := mutex.Unlock(); !ok || err != nil {
		//	logx.Error("redis 解锁失败")
		//	return nil, status.Errorf(codes.Internal, "redis 解锁失败")
		//}
	}

	fmt.Printf("len=%d cap=%d slice=%v\n", len(GoodsDetail), cap(GoodsDetail), GoodsDetail)
	if len := len(GoodsDetail); len > 0 {

		var stockSellDetail = &model.StockSellDetail{
			Detail:  GoodsDetail,
			Status:  1,
			OrderSn: in.OrderSn,
		}

		if result := tx.Create(stockSellDetail); result.RowsAffected == 0 {
			tx.Rollback()
			return nil, status.Errorf(codes.Internal, "保存库存扣减历史失败")
		}
	}

	logx.Infof("回滚了不会执行这边了的吧")
	tx.Commit()

	//m.Unlock()
	return &pb.Empty{}, nil
}
