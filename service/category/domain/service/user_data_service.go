package service

import (
	"github.com/codeleongy/micro-market/service/category/domain/model"
	"github.com/codeleongy/micro-market/service/category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByID(int64) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

func NewCategoryDataService(CategoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{CategoryRepository: CategoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

// 添加分类
func (u *CategoryDataService) AddCategory(Category *model.Category) (CategoryID int64, err error) {
	return u.CategoryRepository.CreateCategory(Category)
}

// 删除分类
func (u *CategoryDataService) DeleteCategory(CategoryID int64) error {
	return u.CategoryRepository.DeleteCategoryByID(CategoryID)
}

// 更新分类
func (u *CategoryDataService) UpdateCategory(Category *model.Category) (err error) {
	return u.CategoryRepository.UpdateCategory(Category)
}

// 根据分类名查找分类
func (u *CategoryDataService) FindCategoryByName(CategoryName string) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByName(CategoryName)
}

// 根据分类ID查找分类
func (u *CategoryDataService) FindCategoryByID(categoryID int64) (*model.Category, error) {
	return u.CategoryRepository.FindCategoryByID(categoryID)
}

// 根据分类分级查找分类
func (u *CategoryDataService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByLevel(level)
}

// 根据分类层级查找分类
func (u *CategoryDataService) FindCategoryByParent(parentID int64) ([]model.Category, error) {
	return u.CategoryRepository.FindCategoryByParent(parentID)
}

// 查找所有分类
func (u *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return u.CategoryRepository.FindAll()
}
