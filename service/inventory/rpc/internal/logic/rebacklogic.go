package logic

import (
	"context"
	"database/sql"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"zero-mal/service/inventory/rpc/internal/svc"
	pb "zero-mal/service/inventory/rpc/inventory_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RebackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRebackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RebackLogic {
	return &RebackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RebackLogic) Reback(in *pb.SellInfo) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	//tx := global.DB.Begin()
	//for _, goods := range in.GoodsInfo {
	//	var inv model.Inventory
	//	if res := global.DB.Where(&model.Inventory{Goods: goods.GoodsId}).First(&inv); res.RowsAffected == 0 {
	//		tx.Rollback()
	//		return nil, status.Errorf(codes.NotFound, "没有库存信息")
	//	}
	//
	//	inv.Stocks += goods.Num
	//	tx.Save(&inv)
	//}
	//tx.Commit()
	//redis 分布式锁
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
	})
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)
	mutex := rs.NewMutex("stock_sell")

	if err := mutex.Lock(); err != nil {
		logx.Error("redis 加锁失败")
	} else {
		logx.Info("redis 加锁成功")
	}
	//if ok, err := mutex.Unlock(); !ok || err != nil {
	//	logx.Error("redis 解锁失败")
	//} else {
	//	logx.Fatal("redis 解锁成功")
	//}

	if err := l.svcCtx.InventoryModel.TransCtx(l.ctx, func(ctx context.Context, conn *gorm.DB) error {

		for _, goods := range in.GoodsInfo {

			inv, err := l.svcCtx.InventoryModel.FindOne(l.ctx, int64(goods.GoodsId))
			if err != nil {
				return status.Errorf(codes.NotFound, "获取库存失败")
			}
			inv.Stocks += goods.Num
			inv.Version += 1 //乐观锁的用法

			if err := l.svcCtx.InventoryModel.Update(l.ctx, inv); err != nil {
				return status.Errorf(codes.InvalidArgument, "修改失败")
			}
			var order sql.NullString

			_ = order.Scan(in.OrderSn)

			Stock, err := l.svcCtx.StockSellDetailModel.FindOneByOrderSn(l.ctx, order)
			if err != nil {
				return status.Errorf(codes.NotFound, "没有出库记录")
			}
			if Stock.Status == 2 {
				return status.Errorf(codes.NotFound, "当前已经回库")
			}
			Stock.Status = 2

			if err := l.svcCtx.StockSellDetailModel.Update(l.ctx, Stock); err != nil {
				return status.Errorf(codes.InvalidArgument, "修改失败")
			}
		}

		return nil
	}); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.Empty{}, nil
}
