package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	model "zero-mal/service/goods/model/gorm"

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

func (l *GetGoodsDetailLogic) GetGoodsDetail(req *pb.GoodInfoRequest) (*pb.GoodsInfoResponse, error) {
	// todo: add your logic here and delete this line
	var good *model.Goods
	good, err := l.svcCtx.GoodsModel.FindOne(l.ctx, int64(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	var brand *model.Brands
	brand, err = l.svcCtx.BrandsModel.FindOne(l.ctx, int64(good.BrandsID))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	good.Brands = *brand

	var category *model.Category
	category, err = l.svcCtx.CategoryModel.FindOne(l.ctx, int64(good.CategoryID))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}
	good.Category = *category

	goodsInfoResponse := ModelToResponse(*good)
	return &goodsInfoResponse, nil

}
