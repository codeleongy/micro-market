package model

type Cart struct {
	// 主键
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	// 商品ID
	ProductID int64 `gorm:"not_null" json:"product_id"`
	// 商品唯一标识
	Num int64 `gorm:"not_null" json:"num"`
	// 尺码ID
	SizeID int64 `gorm:"not_null" json:"size_id"`
	// 用户ID
	UserID int64 `gorm:"not_null" json:"user_id"`
}
