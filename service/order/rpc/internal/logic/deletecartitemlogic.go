package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/service/order/rpc/internal/svc"
	"zero-mal/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCartItemLogic {
	return &DeleteCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCartItemLogic) DeleteCartItem(in *pb.CartItemRequest) (*pb.Empty, error) {
	// todo: add your logic here and delete this line

	if err := l.svcCtx.CartModel.DeleteByUserGood(l.ctx, in.UserId, in.GoodsId); err != nil {
		return nil, status.Errorf(codes.NotFound, "购物车记录不存在")
	}

	return &pb.Empty{}, nil
}
