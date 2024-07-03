package service

import (
	"product/domain/model"
	"product/domain/repository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindAllProduct() ([]model.Product, error)
	FindProductByID(int64) (*model.Product, error)
}

func NewProductDataService(productRepository repository.IProductRepository) IProductDataService {
	return &ProductDataService{ProductRepository: productRepository}
}

type ProductDataService struct {
	ProductRepository repository.IProductRepository
}

// 添加商品
func (u *ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return u.ProductRepository.CreateProduct(product)
}

// 删除商品
func (u *ProductDataService) DeleteProduct(productID int64) error {
	return u.ProductRepository.DeleteProductByID(productID)
}

// 更新商品
func (u *ProductDataService) UpdateProduct(product *model.Product) (err error) {
	return u.ProductRepository.UpdateProduct(product)
}

// 根据商品ID查找商品
func (u *ProductDataService) FindProductByID(productID int64) (*model.Product, error) {
	return u.ProductRepository.FindProductByID(productID)
}

// 查找所有商品
func (u *ProductDataService) FindAllProduct() ([]model.Product, error) {
	return u.ProductRepository.FindAll()
}
