package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/service/order/rpc/internal/svc"
	"zero-mal/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(in *pb.OrderStatus) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	//先查询，再更新 实际上有两条sql执行， select 和 update语句
	if Order, err := l.svcCtx.OrderModel.FindOneByOrderSn(l.ctx, in.OrderSn); err != nil {
		return nil, status.Errorf(codes.NotFound, "订单不存在")
	} else {
		Order.Status = in.Status
		if err := l.svcCtx.OrderModel.Update(l.ctx, Order); err != nil {
			return nil, status.Errorf(codes.NotFound, "订单不存在")
		}
	}

	return &pb.Empty{}, nil
}
