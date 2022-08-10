package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGoodsLogic {
	return &CreateGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGoodsLogic) CreateGoods(in *pb.CreateGoodsInfo) (*pb.GoodsInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GoodsInfoResponse{}, nil
}
