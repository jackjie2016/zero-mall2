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

type UpdateGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGoodsLogic {
	return &UpdateGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGoodsLogic) UpdateGoods(req *pb.CreateGoodsInfo) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	var goods *model.Goods

	goods, err := l.svcCtx.GoodsModel.FindOne(l.ctx, int64(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	var category *model.Category
	//if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
	//	return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	//}
	category, err = l.svcCtx.CategoryModel.FindOne(l.ctx, int64(req.CategoryId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	var brand *model.Brands
	//if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
	//	return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	//}
	brand, err = l.svcCtx.BrandsModel.FindOne(l.ctx, int64(req.BrandId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	goods.Id = req.Id
	//goods.Brands = *brand
	goods.BrandsID = brand.Id
	//goods.Category = *category
	goods.CategoryID = category.Id
	goods.Name = req.Name
	goods.GoodsSn = req.GoodsSn
	goods.MarketPrice = req.MarketPrice
	goods.ShopPrice = req.ShopPrice
	goods.GoodsBrief = req.GoodsBrief
	goods.ShipFree = req.ShipFree
	goods.Images = req.Images
	goods.DescImages = req.DescImages
	goods.GoodsFrontImage = req.GoodsFrontImage
	goods.IsNew = req.IsNew
	goods.IsHot = req.IsHot
	goods.OnSale = req.OnSale

	err = l.svcCtx.GoodsModel.Update(l.ctx, goods)
	if err != nil {
		return nil, err
	}
	//tx := global.DB
	//
	//tx.Begin()
	//
	//result := tx.Save(&goods)
	//
	//if result.Error != nil {
	//	tx.Rollback()
	//	return nil, result.Error
	//}
	//tx.Commit()

	return &pb.Empty{}, nil
}
