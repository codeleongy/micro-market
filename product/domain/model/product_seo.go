package model

type ProductSeo struct {
	ID           int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	SeoTitle     string `json:"seo_title"`
	SeoKeywords  string `json:"seo_keywords"`
	SeoDesc      string `json:"seo_desc"`
	SeoCode      string `gorm:"unique_index;not_null" json:"seo_code"`
	SeoProductID int64  `json:"seo_product_id"`
}
