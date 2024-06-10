package repository

import (
	"github.com/codeleongy/micro-market/service/product/domain/model"
	"github.com/jinzhu/gorm"
)

type IProductRepository interface {
	// 初始化数据表
	InitTable() error
	// 创建商品
	CreateProduct(*model.Product) (int64, error)
	// 根据商品ID删除商品
	DeleteProductByID(int64) error
	// 更新商品信息
	UpdateProduct(*model.Product) error
	// 查找所有商品
	FindAll() ([]model.Product, error)
	// 根据商品ID查找商品信息
	FindProductByID(int64) (*model.Product, error)
}

// 创建UserRepository
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{
		mysqlDB: db,
	}
}

type ProductRepository struct {
	mysqlDB *gorm.DB
}

// 初始化表
func (p *ProductRepository) InitTable() error {
	return p.mysqlDB.CreateTable(
		&model.Product{},
		&model.ProductSeo{},
		&model.ProductImage{},
		&model.ProductSize{},
	).Error
}

// 根据商品ID查找商品信息
func (p *ProductRepository) FindProductByID(productID int64) (*model.Product, error) {
	Product := &model.Product{}
	err := p.mysqlDB.
		Preload("ProductImage").
		Preload("ProductSize").
		Preload("ProductSeo").
		First(Product, productID).Error

	return Product, err
}

// 创建商品
func (p *ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	return product.ID, p.mysqlDB.Create(product).Error
}

// 根据商品ID删除商品
func (p *ProductRepository) DeleteProductByID(productID int64) error {
	// 开启事务
	tx := p.mysqlDB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	// 删除商品
	if err := tx.Unscoped().Where("id = ? ", productID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除商品图片
	if err := tx.Unscoped().Where("images_product_id = ?", productID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除商品规格
	if err := tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除商品SEO
	if err := tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新商品
func (p *ProductRepository) UpdateProduct(product *model.Product) error {
	return p.mysqlDB.Model(product).Update(product).Error
}

// 查找所有类目
func (p *ProductRepository) FindAll() (products []model.Product, err error) {
	err = p.mysqlDB.
		Preload("ProductImage").
		Preload("ProductSeo").
		Preload("ProductSize").
		Find(&products).Error

	return products, err
}
