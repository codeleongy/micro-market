package model

import "time"

type Order struct {
	// 主键
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	// 订单ID
	OrderCode string `gorm:"unique_index;not_null" json:"order_code"`
	// 支付状态
	PayStatus int32 `json:"pay_status"`
	// 发货状态
	ShipStatus int32 `json:"ship_status"`
	// 订单总价
	Price float64 `json:"price"`
	// 订单详情
	OrderDetail []OrderDetail `gorm:"ForeignKey:OrderID" json:"order_detail"`
	// 创建时间
	CreateAt time.Time
	// 更新时间
	UpdateAt time.Time
}
