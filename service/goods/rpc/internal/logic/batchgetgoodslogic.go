package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetGoodsLogic {
	return &BatchGetGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 现在用户提交订单有多个商品，你得批量查询商品的信息吧
func (l *BatchGetGoodsLogic) BatchGetGoods(in *pb.BatchGoodsIdInfo) (*pb.GoodsListResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GoodsListResponse{}, nil
}
