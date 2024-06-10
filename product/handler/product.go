package handler

import (
	"context"

	"github.com/codeleongy/micro-market/product/common"
	"github.com/codeleongy/micro-market/product/domain/model"
	"github.com/codeleongy/micro-market/product/domain/service"
	pb "github.com/codeleongy/micro-market/product/proto/product"
	"go-micro.dev/v4/logger"
)

type Product struct {
	ProductDataService service.IProductDataService
}

func (p *Product) AddProduct(ctx context.Context, req *pb.ProductInfo, res *pb.ResponseProduct) error {
	productAdd := &model.Product{}

	if err := common.SwapTo(req, productAdd); err != nil {
		logger.Error(err)
		return err
	}

	productID, err := p.ProductDataService.AddProduct(productAdd)
	if err != nil {
		logger.Error(err)
		return err
	}

	res.ProductId = productID

	return nil
}

func (p *Product) FindProductByID(ctx context.Context, req *pb.RequestID, res *pb.ProductInfo) error {

	productData, err := p.ProductDataService.FindProductByID(req.ProductId)
	if err != nil {
		return err
	}

	if err := common.SwapTo(productData, res); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (p *Product) UpdateProduct(ctx context.Context, req *pb.ProductInfo, res *pb.Response) error {

	productData := &model.Product{}

	if err := common.SwapTo(req, productData); err != nil {
		logger.Error(err)
		return err
	}

	if err := p.ProductDataService.UpdateProduct(productData); err != nil {
		logger.Error(err)
		return err
	}

	res.Msg = "商品信息更新成功"

	return nil
}

func (p *Product) DeleteProductByID(ctx context.Context, req *pb.RequestID, res *pb.Response) error {
	if err := p.ProductDataService.DeleteProduct(req.ProductId); err != nil {
		logger.Error(err)
		return err
	}

	res.Msg = "商品信息删除成功"

	return nil
}

func (p *Product) FindAllProduct(ctx context.Context, req *pb.RequestAll, res *pb.AllProduct) error {

	products, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		logger.Error(err)
		return err
	}

	for _, product := range products {
		productRes := &pb.ProductInfo{}
		if err := common.SwapTo(product, productRes); err != nil {
			logger.Error(err)
			return err
		}

		res.ProductInfo = append(res.ProductInfo, productRes)
	}

	return nil
}
