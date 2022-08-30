package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"strconv"
	"time"
	"zero-mal/common/tool"

	"github.com/zeromicro/go-zero/core/logx"

	"zero-mal/global"
	"zero-mal/service/goods/rpc/goods_pb"
	"zero-mal/service/inventory/rpc/inventory_pb"
	model "zero-mal/service/order/model/gorm"
	"zero-mal/service/order/rpc/internal/svc"
	"zero-mal/service/order/rpc/pb"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type OrderListener struct {
	Code        codes.Code
	Detail      string
	ID          int32
	OrderAmount float32
	ctx         context.Context
	svx         *svc.ServiceContext
}

//CommitMessageState 状态 会被mq消费者消费
//RollbackMessageState 直接丢弃消息，消费者不会接收到
//UnknowState 未知会被回查
//事务提交，如果是失败或者commit 不回调检查，如果是未知错误就进入checkLocalTransaction
// 是 rollback 还是 commit 根据库存是否扣减，如果库存扣减了，那commit之后，发送一个order_reback 消息
func (o *OrderListener) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {

	var orderInfo model.Order
	var shopCarts []model.Cart

	//parentSpan := opentracing.SpanFromContext(o.ctx) //获取父节点

	_ = json.Unmarshal(msg.Body, &orderInfo)
	//获取购物车产品id
	var goodsIDs []int32

	goodsNumsMap := make(map[int32]int32)
	//shopcartSpan := opentracing.GlobalTracer().StartSpan("shopcart", opentracing.ChildOf(parentSpan.Context()))
	//1. 从购物车中获取到选中的商品
	if res := global.DB.Where(&model.Cart{UserID: orderInfo.UserID, Checked: true}).Find(&shopCarts); res.RowsAffected == 0 {
		o.Code = codes.NotFound
		o.Detail = "购物车中没有商品"
		//shopcartSpan.Finish()
		return primitive.RollbackMessageState
	}
	//shopcartSpan.Finish()

	for _, shopCart := range shopCarts {
		goodsIDs = append(goodsIDs, shopCart.GoodsID)
		goodsNumsMap[shopCart.GoodsID] = shopCart.Nums
	}

	//查询商品微服务获取价格
	//GoodsSrvSpan := opentracing.GlobalTracer().StartSpan("goods-srv", opentracing.ChildOf(parentSpan.Context())), &goods_pb.BatchGoodsIdInfo{Id: goodsIDs}
	goods, err := o.svx.GoodsRpc.BatchGetGoods(context.Background(), &goods_pb.BatchGoodsIdInfo{Id: goodsIDs})
	if err != nil {
		o.Code = codes.Internal
		o.Detail = "批量查询商品信息失败"
		return primitive.RollbackMessageState
	}
	//GoodsSrvSpan.Finish()

	//计算库存
	var orderAmount float32
	var orderGoods []*model.OrderGoods
	var goodsInvInfo []*inventory_pb.GoodsInvInfo
	for _, good := range goods.Data {
		orderAmount += float32(goodsNumsMap[good.Id]) * good.ShopPrice
		orderGoods = append(orderGoods, &model.OrderGoods{
			GoodsID:    good.Id,
			GoodsName:  good.Name,
			GoodsImage: good.GoodsFrontImage,
			GoodsPrice: good.ShopPrice,
			Nums:       goodsNumsMap[good.Id],
		})
		goodsInvInfo = append(goodsInvInfo, &inventory_pb.GoodsInvInfo{
			GoodsId: good.Id,
			Num:     goodsNumsMap[good.Id],
		})
	}

	//跨服务调用库存微服务进行库存扣减
	//queryInvSpan := opentracing.GlobalTracer().StartSpan("inv-srv", opentracing.ChildOf(parentSpan.Context()))
	if _, err = o.svx.InventoryRpc.Sell(context.Background(), &inventory_pb.SellInfo{GoodsInfo: goodsInvInfo, OrderSn: orderInfo.OrderSn}); err != nil {
		//需要深入判断是什么原因导致的失败是网络问题还是程序问题，或者数据问题
		logx.Errorf("扣减库存失败:%v", err.Error())
		o.Code = codes.ResourceExhausted
		o.Detail = "扣减库存失败"
		return primitive.RollbackMessageState
		//return nil, status.Errorf(codes.ResourceExhausted, "扣减库存失败")
	}
	//queryInvSpan.Finish()

	//生成订单表
	//20210308xxxx
	tx := global.DB.Begin()
	orderInfo.OrderMount = orderAmount
	//SaveOrderSpan := opentracing.GlobalTracer().StartSpan("save-order", opentracing.ChildOf(parentSpan.Context()))
	if result := tx.Save(&orderInfo); result.RowsAffected == 0 {
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "创建订单失败"
		logx.Errorf("创建订单失败:%v", err.Error())
		return primitive.CommitMessageState //创建失败立即回滚
	}
	//SaveOrderSpan.Finish()

	for _, orderGood := range orderGoods {
		orderGood.Order = orderInfo.Id
	}

	//批量插入orderGoods
	//SaveGoodsSpan := opentracing.GlobalTracer().StartSpan("save-goods", opentracing.ChildOf(parentSpan.Context()))
	if result := tx.CreateInBatches(orderGoods, 100); result.RowsAffected == 0 {
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "创建订单商品明细失败"
		logx.Errorf("创建订单商品明细失败:%v", err)
		return primitive.CommitMessageState //创建失败立即回滚

	}
	//SaveGoodsSpan.Finish()

	//DeleteCartSpan := opentracing.GlobalTracer().StartSpan("delete-cart", opentracing.ChildOf(parentSpan.Context()))
	if result := tx.Where(&model.Cart{UserID: orderInfo.UserID, Checked: true}).Delete(&model.Cart{}); result.RowsAffected == 0 {
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "删除购物车失败"
		logx.Errorf("删除购物车失败:%v", err.Error())
		return primitive.CommitMessageState //创建失败立即回滚

	}
	//DeleteCartSpan.Finish()

	o.OrderAmount = orderAmount
	o.ID = orderInfo.Id
	o.Code = codes.OK

	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(o.svx.Config.Rocketmq.Hosts)),
		producer.WithRetry(2),
	)

	if err != nil {
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "生成延迟消息producer失败"
		logx.Errorf("生成 producer 失败: %s", err.Error())
		return primitive.CommitMessageState //创建失败立即回滚
	}

	err = p.Start()
	if err != nil {
		//os.Exit(1)
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "启动延迟消息producer失败"
		logx.Errorf("启动延迟消息producer失败: %s", err.Error())
		return primitive.CommitMessageState //创建失败立即回滚
	}

	msg2 := &primitive.Message{
		Topic: "order_timeout",
		Body:  msg.Body,
	}
	msg2.WithDelayTimeLevel(4) //跟普通比就多一句这个
	_, err = p.SendSync(context.Background(), msg2)
	if err != nil {
		//os.Exit(1)
		tx.Rollback()
		o.Code = codes.Internal
		o.Detail = "发送延迟消息失败"
		logx.Errorf("发送延迟消息失败: %s", err.Error())
		return primitive.CommitMessageState //创建失败立即回滚
	}
	//提交事务
	tx.Commit()
	return primitive.RollbackMessageState
}

// When no response to prepare(half) message. broker will send check message to check the transaction status, and this
// method will be invoked to get local transaction status.
// 回调检查
func (o *OrderListener) CheckLocalTransaction(*primitive.MessageExt) primitive.LocalTransactionState {

	return primitive.UnknowState
}
func GenerateOrderSn(userId int32) string {
	//订单号的生成规则
	/*
		年月日时分秒+用户id+2位随机数
	*/
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(),
		userId, rand.Intn(90)+10,
	)
	return orderSn
}

// 订单
func (l *CreateOrderLogic) CreateOrder(in *pb.OrderRequest) (*pb.OrderInfoResponse, error) {
	// todo: add your logic here and delete this line

	/*
		新建订单
			1. 从购物车中获取到选中的商品
			2. 商品的价格自己查询 - 访问商品服务 (跨微服务)
			3. 库存的扣减 - 访问库存服务 (跨微服务)
			4. 订单的基本信息表 - 订单的商品信息表
			5. 从购物车中删除已购买的记录
	*/

	orderListener := OrderListener{ctx: l.ctx, svx: l.svcCtx}
	p, err := rocketmq.NewTransactionProducer(
		&orderListener,
		producer.WithNsResolver(primitive.NewPassthroughResolver(l.svcCtx.Config.Rocketmq.Hosts)),
		producer.WithRetry(1),
	)

	if err != nil {
		logx.Errorf("生成 producer 失败: %s", err.Error())
		return nil, err
	}

	err = p.Start()

	if err != nil {
		logx.Errorf("启动 producer 失败: %s", err.Error())
		//os.Exit(1)
		return nil, err
	}

	topic := "order_reback"
	worker, _ := tool.NewOrderWorker(0)
	order := model.Order{
		OrderSn:      strconv.Itoa(int(worker.NextOrderId())),
		Address:      in.Address,
		SignerName:   in.Name,
		SingerMobile: in.Mobile,
		Post:         in.Post,
		UserID:       in.UserId,
		Status:       "PAYING",
	}
	jsonString, _ := json.Marshal(order)
	msg := &primitive.Message{
		Topic: topic,
		Body:  jsonString,
	}
	_, err = p.SendMessageInTransaction(context.Background(), msg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "发送消息失败")
	}

	if orderListener.Code != codes.OK {
		return nil, status.Errorf(orderListener.Code, orderListener.Detail)
	}

	return &pb.OrderInfoResponse{Id: orderListener.ID, OrderSn: order.OrderSn, Total: orderListener.OrderAmount}, nil
}
