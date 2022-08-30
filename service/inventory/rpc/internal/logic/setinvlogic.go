package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/service/inventory/rpc/internal/svc"
	pb "zero-mal/service/inventory/rpc/inventory_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetInvLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetInvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetInvLogic {
	return &SetInvLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetInvLogic) SetInv(in *pb.GoodsInvInfo) (*pb.Empty, error) {
	// todo: add your logic here and delete this line

	inv, err := l.svcCtx.InventoryModel.FindOne(l.ctx, int64(in.GoodsId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "获取库存异常")
	}
	fmt.Println(inv)

	inv.GoodsID = in.GoodsId
	inv.Stocks = in.Num

	err = l.svcCtx.InventoryModel.Update(l.ctx, inv)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "库存设置失败")
	}
	return &pb.Empty{}, nil
}
