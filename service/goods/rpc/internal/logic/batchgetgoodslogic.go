package logic

import (
	"context"
	"zero-mal/global"
	model "zero-mal/service/goods/model/gorm"

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
func (l *BatchGetGoodsLogic) BatchGetGoods(req *pb.BatchGoodsIdInfo) (*pb.GoodsListResponse, error) {
	// todo: add your logic here and delete this line
	goodsListResponse := &pb.GoodsListResponse{}
	var goods []model.Goods

	//调用where并不会真正执行sql 只是用来生成sql的 当调用find， first才会去执行sql，
	result := global.DB.Preload("Category").Preload("Brands").Where(req.Id).Find(&goods)
	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	goodsListResponse.Total = int32(result.RowsAffected)
	return goodsListResponse, nil

}
