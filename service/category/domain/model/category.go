package model

type Category struct {
	// 主键
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	// 分类名称
	CategoryName string `gorm:"unique_index;not_null" json:"category_name"`
	// 分类级别
	CategoryLevel uint32 `json:"category_level"`
	// 分类父级
	CategoryParent int64 `json:"category_parent"`
	// 分类图片
	CategoryImage string `json:"category_image"`
	// 分类描述
	CategoryDescription string `json:"category_description"`
}
