package model

type Payment struct {
	// 主键
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	// 支付名称
	PaymentName string `json:"payment_name"`
	// 支付唯一ID
	PaymentSid string `json:"payment_sid"`
	// 支付通道状态 沙盒（false） 生产（true）
	PaymentStatus bool `json:"payment_status"`
	// 支付图片或者logo
	PaymentImage float64 `json:"payment_image"`
}
