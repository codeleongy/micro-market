package handler

import (
	"context"
	"micro-market/common"
	"order/domain/model"
	"order/domain/service"
	pb "order/proto"

	"go-micro.dev/v4/logger"
)

type Order struct {
	OrderDataService service.IOrderDataService
}

func (o *Order) GetOrderByID(ctx context.Context, req *pb.OrderID, rsp *pb.OrderInfo) error {
	orderInfo, err := o.OrderDataService.FindOrderByID(req.OrderId)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = common.SwapTo(orderInfo, rsp)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (o *Order) GetAllOrder(ctx context.Context, req *pb.AllOrderReq, rsp *pb.AllOrder) error {
	orders, err := o.OrderDataService.FindAllOrder()
	if err != nil {
		logger.Error(err)
		return err
	}

	for _, v := range orders {
		orderInfo := &pb.OrderInfo{}
		err := common.SwapTo(v, orderInfo)
		if err != nil {
			logger.Error(err)
			return err
		}
		rsp.OrderInfos = append(rsp.OrderInfos, orderInfo)
	}
	return nil
}

func (o *Order) CreateOrder(ctx context.Context, req *pb.OrderInfo, rsp *pb.OrderID) error {
	orderInfo := &model.Order{}

	err := common.SwapTo(req, orderInfo)
	if err != nil {
		logger.Error(err)
		return err
	}

	orderId, err := o.OrderDataService.AddOrder(orderInfo)
	if err != nil {
		logger.Error(err)
		return err
	}

	rsp.OrderId = orderId
	return nil
}

func (o *Order) DeleteOrderByID(ctx context.Context, req *pb.OrderID, rsp *pb.Response) error {
	if err := o.OrderDataService.DeleteOrder(req.OrderId); err != nil {
		logger.Error(err)
		return err
	}

	rsp.Msg = "订单删除成功"
	return nil
}

func (o *Order) UpdateOrderPayStatus(ctx context.Context, req *pb.PayStatus, rsp *pb.Response) error {
	err := o.OrderDataService.UpdatePayStatus(req.OrderId, req.PayStatus)
	if err != nil {
		logger.Error(err)
		return err
	}

	rsp.Msg = "订单支付状态更新成功"
	return nil
}

// 更新发货状态
func (o *Order) UpdateOrderShipStatus(ctx context.Context, req *pb.ShipStatus, rsp *pb.Response) error {

	err := o.OrderDataService.UpdateShipStatus(req.OrderId, req.ShipStatus)
	if err != nil {
		logger.Error(err)
		return err
	}

	rsp.Msg = "订单发货状态更新成功"

	return nil

}
func (o *Order) UpdateOrder(ctx context.Context, req *pb.OrderInfo, rsp *pb.Response) error {

	orderInfo := &model.Order{}

	err := common.SwapTo(req, orderInfo)
	if err != nil {
		logger.Error(err)
		return err
	}

	err = o.OrderDataService.UpdateOrder(orderInfo)
	if err != nil {
		logger.Error(err)
		return err
	}

	rsp.Msg = "订单更新成功"
	return nil

}
