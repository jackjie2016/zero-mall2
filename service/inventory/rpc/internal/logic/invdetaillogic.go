package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"zero-mal/service/inventory/rpc/internal/svc"
	pb "zero-mal/service/inventory/rpc/inventory_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvDetailLogic {
	return &InvDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvDetailLogic) InvDetail(in *pb.GoodsInvInfo) (*pb.GoodsInvInfo, error) {
	// todo: add your logic here and delete this line

	Inventory, err := l.svcCtx.InventoryModel.FindOne(l.ctx, int64(in.GoodsId))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "获取库存异常")
	}

	return &pb.GoodsInvInfo{
		GoodsId: Inventory.GoodsID,
		Num:     Inventory.Stocks,
	}, nil
}
