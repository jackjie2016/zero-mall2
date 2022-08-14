package logic

import (
	"context"
	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGoodsLogic {
	return &DeleteGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGoodsLogic) DeleteGoods(req *pb.DeleteGoodsInfo) (*pb.Empty, error) {
	// todo: add your logic here and delete this line

	_, err := l.svcCtx.GoodsModel.FindOne(l.ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.GoodsModel.Delete(l.ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
