package repository

import (
	"errors"

	"github.com/codeleongy/micro-market/cart/domain/model"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	// 初始化数据表
	InitTable() error
	// 创建购物车
	CreateCart(*model.Cart) (int64, error)
	// 根据购物车ID删除购物车
	DeleteCartByID(int64) error
	// 更新购物车信息
	UpdateCart(*model.Cart) error
	// 查找所有购物车
	FindAll(int64) ([]model.Cart, error)
	// 根据购物车ID查找购物车信息
	FindCartByID(int64) (*model.Cart, error)

	// 清空购物车
	CleanCart(int64) error
	// 添加购物车物品
	IncrNum(int64, int64) error
	// 减少购物车物品
	DecrNum(int64, int64) error
}

// 创建UserRepository
func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{
		mysqlDB: db,
	}
}

type CartRepository struct {
	mysqlDB *gorm.DB
}

// 初始化表
func (c *CartRepository) InitTable() error {
	return c.mysqlDB.CreateTable(&model.Cart{}).Error
}

// 根据购物车ID查找购物车信息
func (c *CartRepository) FindCartByID(CartID int64) (*model.Cart, error) {
	Cart := &model.Cart{}
	err := c.mysqlDB.First(Cart, CartID).Error

	return Cart, err
}

// 创建购物车
func (c *CartRepository) CreateCart(cart *model.Cart) (int64, error) {
	db := c.mysqlDB.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.ProductID, UserID: cart.UserID})
	if db.Error != nil {
		return 0, db.Error
	}

	if db.RowsAffected == 0 {
		return 0, errors.New("购物车创建失败")
	}

	return cart.ID, nil
}

// 根据购物车ID删除购物车
func (c *CartRepository) DeleteCartByID(cartID int64) error {
	return c.mysqlDB.Where("id = ?", cartID).Delete(&model.Cart{}).Error
}

// 更新购物车
func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDB.Model(cart).Update(cart).Error
}

// 查找所有类目
func (c *CartRepository) FindAll(userID int64) (carts []model.Cart, err error) {
	err = c.mysqlDB.
		Where("user_id = ?", userID).
		Find(&carts).Error
	return carts, err
}

// 根据用户ID清空购物车
func (c *CartRepository) CleanCart(userID int64) error {
	return c.mysqlDB.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}

// 添加商品数量
func (c *CartRepository) IncrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	return c.mysqlDB.Model(cart).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

// 减少商品物品
func (c *CartRepository) DecrNum(cartID int64, num int64) error {
	cart := &model.Cart{ID: cartID}
	db := c.mysqlDB.Model(cart).Where("num >= ?", num).UpdateColumn("num", gorm.Expr("num - ?", num))

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("减少失败")
	}
	return nil
}
