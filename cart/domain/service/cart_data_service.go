package service

import (
	"github.com/codeleongy/micro-market/cart/domain/model"
	"github.com/codeleongy/micro-market/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindAllCart(int64) ([]model.Cart, error)
	FindCartByID(int64) (*model.Cart, error)
	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartDataService(CartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{CartRepository: CartRepository}
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

// 添加商品
func (c *CartDataService) AddCart(Cart *model.Cart) (int64, error) {
	return c.CartRepository.CreateCart(Cart)
}

// 删除商品
func (c *CartDataService) DeleteCart(CartID int64) error {
	return c.CartRepository.DeleteCartByID(CartID)
}

// 更新商品
func (c *CartDataService) UpdateCart(Cart *model.Cart) (err error) {
	return c.CartRepository.UpdateCart(Cart)
}

// 根据商品ID查找商品
func (c *CartDataService) FindCartByID(CartID int64) (*model.Cart, error) {
	return c.CartRepository.FindCartByID(CartID)
}

// 查找所有商品
func (c *CartDataService) FindAllCart(userID int64) ([]model.Cart, error) {
	return c.CartRepository.FindAll(userID)
}

// 清空购物车商品
func (c *CartDataService) CleanCart(userID int64) error {
	return c.CartRepository.CleanCart(userID)
}

// 增加购物车商品
func (c *CartDataService) IncrNum(cartID int64, num int64) error {
	return c.CartRepository.IncrNum(cartID, num)
}

// 减少购物车商品
func (c *CartDataService) DecrNum(cartID int64, num int64) error {
	return c.CartRepository.DecrNum(cartID, num)
}
