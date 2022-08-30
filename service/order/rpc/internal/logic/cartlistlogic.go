package logic

import (
	"context"
	"zero-mal/global"
	model "zero-mal/service/order/model/gorm"

	"zero-mal/service/order/rpc/internal/svc"
	"zero-mal/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartListLogic {
	return &CartListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 购物车
func (l *CartListLogic) CartList(in *pb.UserInfo) (*pb.CartListResponse, error) {
	// todo: add your logic here and delete this line
	var shopCarts []model.Cart
	var rsp pb.CartListResponse

	if res := global.DB.Where(&model.Cart{UserID: in.Id}).Find(&shopCarts); res.Error != nil {
		return nil, res.Error
	} else {
		rsp.Total = int32(res.RowsAffected)
	}

	for _, shopCart := range shopCarts {
		rsp.Data = append(rsp.Data, &pb.ShopCartInfoResponse{
			Id:      shopCart.Id,
			UserId:  shopCart.UserID,
			GoodsId: shopCart.GoodsID,
			Nums:    shopCart.Nums,
			Checked: shopCart.Checked,
		})
	}
	return &rsp, nil
}
