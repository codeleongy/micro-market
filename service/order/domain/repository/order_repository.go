package repository

import (
	"errors"
	"order/domain/model"

	"github.com/jinzhu/gorm"
)

type IOrderRepository interface {
	// 初始化数据表
	InitTable() error
	// 创建订单
	CreateOrder(*model.Order) (int64, error)
	// 根据订单ID删除订单
	DeleteOrderByID(int64) error
	// 更新订单信息
	UpdateOrder(*model.Order) error
	// 查找所有订单
	FindAll() ([]model.Order, error)
	// 根据订单ID查找订单信息
	FindOrderByID(int64) (*model.Order, error)
	// 更新发货状态
	UpdateShipStatus(int64, int32) error
	// 更新支付状态
	UpdatePayStatus(int64, int32) error
}

// 创建OrderRepository
func NewOrderRepository(db *gorm.DB) IOrderRepository {
	return &OrderRepository{
		mysqlDB: db,
	}
}

type OrderRepository struct {
	mysqlDB *gorm.DB
}

// 初始化表
func (r *OrderRepository) InitTable() error {
	return r.mysqlDB.CreateTable(&model.Order{}, &model.OrderDetail{}).Error
}

// 根据订单ID查找订单信息
func (r *OrderRepository) FindOrderByID(orderID int64) (*model.Order, error) {
	Order := &model.Order{}
	err := r.mysqlDB.Preload("OrderDetail").First(Order, orderID).Error

	return Order, err
}

// 创建订单
func (r *OrderRepository) CreateOrder(order *model.Order) (int64, error) {
	return order.ID, r.mysqlDB.Create(order).Error
}

// 根据订单ID删除订单
func (r *OrderRepository) DeleteOrderByID(orderID int64) error {
	tx := r.mysqlDB.Begin()
	// 遇到错误回滚
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// 彻底删除Order信息
	if err := tx.Unscoped().Where("id = ?", orderID).Delete(&model.Order{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 彻底删除OrderDetail信息
	if err := tx.Unscoped().Where("order_id = ?", orderID).Delete(&model.OrderDetail{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新订单
func (r *OrderRepository) UpdateOrder(order *model.Order) error {
	return r.mysqlDB.Model(order).Update(order).Error
}

// 查找所有类目
func (r *OrderRepository) FindAll() (orders []model.Order, err error) {
	err = r.mysqlDB.
		Preload("OrderDetail").
		Find(&orders).Error
	return orders, err
}

// 更新发货状态
func (r *OrderRepository) UpdateShipStatus(orderID int64, shipStatus int32) error {
	db := r.mysqlDB.Model(&model.Order{}).Where("id = ?", orderID).UpdateColumn("ship_status", shipStatus)

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("发货状态更新失败")
	}

	return nil
}

// 更新支付状态
func (r *OrderRepository) UpdatePayStatus(orderId int64, payStatus int32) error {
	db := r.mysqlDB.Model(&model.Order{}).Where("id = ?", orderId).UpdateColumn("pay_status", payStatus)

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("支付状态更新失败")
	}

	return nil
}
