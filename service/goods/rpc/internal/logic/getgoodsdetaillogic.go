package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsDetailLogic {
	return &GetGoodsDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsDetailLogic) GetGoodsDetail(in *pb.GoodInfoRequest) (*pb.GoodsInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GoodsInfoResponse{}, nil
}
