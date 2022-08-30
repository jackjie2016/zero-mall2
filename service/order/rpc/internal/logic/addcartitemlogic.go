package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	model "zero-mal/service/order/model/gorm"

	"zero-mal/service/order/rpc/internal/svc"
	"zero-mal/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartItemLogic {
	return &AddCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCartItemLogic) AddCartItem(in *pb.CartItemRequest) (*pb.ShopCartInfoResponse, error) {
	// todo: add your logic here and delete this line

	//将商品添加到购物车 1. 购物车中原本没有这件商品 - 新建一个记录 2. 这个商品之前添加到了购物车- 合并
	var shopCart *model.Cart
	var err error

	if shopCart, err = l.svcCtx.CartModel.FindOne(l.ctx, int64(in.GoodsId)); err != nil {
		return nil, status.Errorf(codes.NotFound, "购物车记录不存在")
	}
	if shopCart != nil {

		shopCart.Nums += in.Nums

		if err = l.svcCtx.CartModel.Update(l.ctx, shopCart); err != nil {
			return nil, status.Errorf(codes.Internal, "购物车更新失败")
		}

	} else {
		shopCart.UserID = in.UserId
		shopCart.GoodsID = in.GoodsId
		shopCart.Nums = in.Nums
		shopCart.Checked = false

		if err = l.svcCtx.CartModel.Insert(l.ctx, shopCart); err != nil {
			return nil, status.Errorf(codes.Internal, "购物车添加失败")
		}
	}

	return &pb.ShopCartInfoResponse{Id: int32(shopCart.Id)}, nil

}
