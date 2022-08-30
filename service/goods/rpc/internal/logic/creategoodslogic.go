package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/global"
	model "zero-mal/service/goods/model/gorm"

	"zero-mal/service/goods/rpc/goods_pb"
	"zero-mal/service/goods/rpc/internal/svc"

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

func (l *CreateGoodsLogic) CreateGoods(req *goods_pb.CreateGoodsInfo) (*goods_pb.GoodsInfoResponse, error) {
	// todo: add your logic here and delete this line

	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	//这里没有看到图片文件是如何上传， 在微服务中 普通的文件上传已经不再使用
	goods := model.Goods{
		Brands:          brand,
		BrandsID:        brand.Id,
		Category:        category,
		CategoryID:      category.Id,
		Name:            req.Name,
		GoodsSn:         req.GoodsSn,
		MarketPrice:     req.MarketPrice,
		ShopPrice:       req.ShopPrice,
		GoodsBrief:      req.GoodsBrief,
		ShipFree:        req.ShipFree,
		Images:          req.Images,
		DescImages:      req.DescImages,
		GoodsFrontImage: req.GoodsFrontImage,
		IsNew:           req.IsNew,
		IsHot:           req.IsHot,
		OnSale:          req.OnSale,
	}

	//srv之间互相调用了
	tx := global.DB
	tx.Begin()

	result := tx.Save(&goods)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()

	return &goods_pb.GoodsInfoResponse{
		Id: goods.Id,
	}, nil
}
