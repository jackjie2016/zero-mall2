package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zero-mal/global"
	"zero-mal/service/goods/rpc/goods"
	"zero-mal/service/inventory/rpc/inventory"
	model "zero-mal/service/order/model/gorm"
	"zero-mal/service/order/rpc/internal/config"
	"zero-mal/service/order/rpc/internal/initialize"
)

type ServiceContext struct {
	Config config.Config

	CartModel       model.CartModel
	OrderModel      model.OrderModel
	OrderGoodsModel model.OrderGoodsModel

	//UserRpc      usercenter.Usercenter
	GoodsRpc     goods.Goods
	InventoryRpc inventory.Inventory
}

func NewServiceContext(c config.Config) *ServiceContext {
	initialize.InitDb(c.DB.DataSource)

	return &ServiceContext{
		Config:          c,
		CartModel:       model.NewCartModel(global.DB, c.Cache),
		OrderModel:      model.NewOrderModel(global.DB, c.Cache),
		OrderGoodsModel: model.NewOrderGoodsModel(global.DB, c.Cache),

		//UserRpc:      usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpc)),
		GoodsRpc:     goods.NewGoods(zrpc.MustNewClient(c.GoodsRpc)),
		InventoryRpc: inventory.NewInventory(zrpc.MustNewClient(c.InventoryRpc)),
	}
}
