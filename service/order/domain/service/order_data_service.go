package service

import (
	"order/domain/model"
	"order/domain/repository"
)

type IOrderDataService interface {
	AddOrder(*model.Order) (int64, error)
	DeleteOrder(int64) error
	UpdateOrder(*model.Order) error
	FindAllOrder() ([]model.Order, error)
	FindOrderByID(int64) (*model.Order, error)
	UpdateShipStatus(int64, int32) error
	UpdatePayStatus(int64, int32) error
}

func NewOrderDataService(orderRepository repository.IOrderRepository) IOrderDataService {
	return &OrderDataService{OrderRepository: orderRepository}
}

type OrderDataService struct {
	OrderRepository repository.IOrderRepository
}

// 添加订单
func (o *OrderDataService) AddOrder(order *model.Order) (int64, error) {
	return o.OrderRepository.CreateOrder(order)
}

// 删除订单
func (o *OrderDataService) DeleteOrder(OrderID int64) error {
	return o.OrderRepository.DeleteOrderByID(OrderID)
}

// 更新订单
func (o *OrderDataService) UpdateOrder(Order *model.Order) (err error) {
	return o.OrderRepository.UpdateOrder(Order)
}

// 根据订单ID查找订单
func (o *OrderDataService) FindOrderByID(OrderID int64) (*model.Order, error) {
	return o.OrderRepository.FindOrderByID(OrderID)
}

// 查找所有订单
func (o *OrderDataService) FindAllOrder() ([]model.Order, error) {
	return o.OrderRepository.FindAll()
}

func (o *OrderDataService) UpdateShipStatus(orderId int64, shipStatus int32) error {
	return o.OrderRepository.UpdateShipStatus(orderId, shipStatus)
}
func (o *OrderDataService) UpdatePayStatus(orderId int64, payStatus int32) error {
	return o.OrderRepository.UpdatePayStatus(orderId, payStatus)
}
