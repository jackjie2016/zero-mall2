package logic

import (
	"context"
	"zero-mal/common/tool"
	"zero-mal/global"

	model "zero-mal/service/order/model/gorm"
	"zero-mal/service/order/rpc/internal/svc"
	"zero-mal/service/order/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderListLogic) OrderList(in *pb.OrderFilterRequest) (*pb.OrderListResponse, error) {
	// todo: add your logic here and delete this line
	var orders []model.Order
	var rsp pb.OrderListResponse

	var total int64
	global.DB.Where(&model.Order{UserID: in.UserId}).Count(&total)
	rsp.Total = int32(total)

	//分页
	global.DB.Scopes(tool.Paginate(int(in.Pages), int(in.PagePerNums))).Where(&model.Order{UserID: in.UserId}).Find(&orders)
	for _, order := range orders {
		rsp.Data = append(rsp.Data, &pb.OrderInfoResponse{
			Id:      order.Id,
			UserId:  order.UserID,
			OrderSn: order.OrderSn,
			PayType: order.PayType,
			Status:  order.Status,
			Post:    order.Post,
			Total:   order.OrderMount,
			Address: order.Address,
			Name:    order.SignerName,
			Mobile:  order.SingerMobile,
			AddTime: order.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &rsp, nil

}
