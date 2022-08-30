package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/global"
	model "zero-mal/service/order/model/gorm"

	"zero-mal/service/order/rpc/internal/svc"
	"zero-mal/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemLogic {
	return &UpdateCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCartItemLogic) UpdateCartItem(in *pb.CartItemRequest) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	//更新或者插入

	var shopCart model.Cart

	if result := global.DB.Where(&model.Cart{UserID: in.UserId, GoodsID: in.GoodsId}).First(&shopCart); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "购物车记录不存在")
	}
	shopCart.Checked = in.Checked
	if in.Nums > 0 {
		shopCart.Nums = in.Nums
	}

	global.DB.Save(&shopCart)

	return &pb.Empty{}, nil
}
