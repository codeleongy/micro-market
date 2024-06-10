package handler

import (
	"context"

	"github.com/codeleongy/micro-market/category/domain/model"
	"github.com/codeleongy/micro-market/category/domain/service"
	pb "github.com/codeleongy/micro-market/category/proto/category"
	"github.com/codeleongy/micro-market/common"
	"go-micro.dev/v4/logger"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

// 分类创建
func (c *Category) CreateCategory(ctx context.Context, req *pb.CategoryReq, res *pb.CreateCategoryRes) error {
	category := &model.Category{}
	// 赋值
	err := common.SwapTo(req, category)
	if err != nil {
		return err
	}
	categoryID, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}

	res.CategoryId = categoryID
	res.Message = "分类添加成功"

	return nil
}

func (c *Category) UpdateCategory(ctx context.Context, req *pb.CategoryReq, res *pb.UpdateCategoryRes) error {
	category := &model.Category{}

	err := common.SwapTo(req, category)
	if err != nil {
		return err
	}

	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}

	res.Message = "分类更新成功"

	return nil
}

func (c *Category) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryReq, res *pb.DeleteCategoryRes) error {
	err := c.CategoryDataService.DeleteCategory(req.CategoryId)
	if err != nil {
		return err
	}

	res.Message = "删除成功"
	return nil
}

func (c *Category) FindCategoryByName(ctx context.Context, req *pb.FindByNameReq, res *pb.CategoryRes) error {
	category, err := c.CategoryDataService.FindCategoryByName(req.CategoryName)
	if err != nil {
		return err
	}

	return common.SwapTo(category, res)
}

func (c *Category) FindCategoryByID(ctx context.Context, req *pb.FindByIDReq, res *pb.CategoryRes) error {
	category, err := c.CategoryDataService.FindCategoryByID(req.CategoryId)
	if err != nil {
		return err
	}

	return common.SwapTo(category, res)

}
func (c *Category) FindCategoryByLevel(ctx context.Context, req *pb.FindByLevelReq, res *pb.FindAllRes) error {
	category, err := c.CategoryDataService.FindCategoryByLevel(req.Level)
	if err != nil {
		return err
	}

	return common.SwapTo(category, res)
}

func (c *Category) FindCategoryByParent(ctx context.Context, req *pb.FindByParentReq, res *pb.FindAllRes) error {
	categorys, err := c.CategoryDataService.FindCategoryByParent(req.ParentId)
	if err != nil {
		return err
	}

	for _, category := range categorys {
		cr := &pb.CategoryRes{}
		common.SwapTo(category, cr)
		if err != nil {
			logger.Error(err)
			break
		}
		res.Categorys = append(res.Categorys, cr)
	}

	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, req *pb.FindAllReq, res *pb.FindAllRes) error {
	categorys, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}

	for _, category := range categorys {
		cr := &pb.CategoryRes{}
		common.SwapTo(category, cr)
		if err != nil {
			logger.Error(err)
			break
		}
		res.Categorys = append(res.Categorys, cr)
	}

	return nil
}
