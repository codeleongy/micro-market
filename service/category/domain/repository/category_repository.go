package repository

import (
	"category/domain/model"

	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	// 初始化数据表
	InitTable() error
	// 创建分类
	CreateCategory(*model.Category) (int64, error)
	// 根据分类ID删除分类
	DeleteCategoryByID(int64) error
	// 更新分类信息
	UpdateCategory(*model.Category) error
	// 查找所有分类
	FindAll() ([]model.Category, error)
	// 根据分类名称查找用信息
	FindCategoryByName(string) (*model.Category, error)
	// 根据分类ID查找分类信息
	FindCategoryByID(int64) (*model.Category, error)
	// 根据
	FindCategoryByLevel(uint32) ([]model.Category, error)
	// xxx
	FindCategoryByParent(int64) ([]model.Category, error)
}

// 创建UserRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{
		mysqlDB: db,
	}
}

type CategoryRepository struct {
	mysqlDB *gorm.DB
}

// 初始化表
func (c *CategoryRepository) InitTable() error {
	return c.mysqlDB.CreateTable(&model.Category{}).Error
}

// 根据分类名称查找用信息
func (c *CategoryRepository) FindCategoryByName(name string) (*model.Category, error) {
	user := &model.Category{}
	return user, c.mysqlDB.Where("category_name = ?", name).Find(user).Error
}

func (c *CategoryRepository) FindCategoryByLevel(level uint32) (categorys []model.Category, err error) {
	return categorys, c.mysqlDB.Where("category_level = ?", level).Find(categorys).Error
}

func (c *CategoryRepository) FindCategoryByParent(parent int64) (categorys []model.Category, err error) {
	return categorys, c.mysqlDB.Where("category_parent = ?", parent).Find(&categorys).Error
}

// 根据分类ID查找分类信息
func (c *CategoryRepository) FindCategoryByID(categoryID int64) (*model.Category, error) {
	category := &model.Category{}
	return category, c.mysqlDB.First(category, categoryID).Error
}

// 创建分类
func (c *CategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	return category.ID, c.mysqlDB.Create(category).Error
}

// 根据分类ID删除分类
func (c *CategoryRepository) DeleteCategoryByID(categoryID int64) error {
	return c.mysqlDB.Where("id = ?", categoryID).Delete(&model.Category{}).Error
}

// 更新分类
func (c *CategoryRepository) UpdateCategory(category *model.Category) error {
	return c.mysqlDB.Model(category).Update(category).Error
}

// 查找所有类目
func (c *CategoryRepository) FindAll() (categorys []model.Category, err error) {
	return categorys, c.mysqlDB.Find(categorys).Error
}
