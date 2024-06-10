package model

type Product struct {
	// 主键
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品唯一标识
	ProductSku string `gorm:"unique_index;not_null" json:"product_sku"`
	// 商品价格
	ProductPrice float64 `json:"product_price"`
	// 商品描述
	ProductDesc string `json:"product_desc"`
	// 商品图片
	ProductImage []ProductImage `gorm:"ForeignKey:ImageProductID" json:"product_image"`
	// 商品尺码
	ProductSize []ProductSize `gorm:"ForeignKey:SizeProductID" json:"product_size"`
	// SEO
	ProductSeo ProductSeo `gorm:"ForeignKey:SeoProductID" json:"product_seo"`
}
