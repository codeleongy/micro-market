package main

import (
	"context"
	"fmt"

	"micro-market/common"
	pb "product/proto"

	"github.com/go-micro/plugins/v4/registry/consul"
	opentracingFn "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

func main() {

	// 注册中心
	consulRegistry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"127.0.0.1:18500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.product.client", "localhost:6831")
	if err != nil {
		logger.Error(err)
	}

	defer io.Close()

	opentracing.SetGlobalTracer(t)

	srv := micro.NewService()

	srv.Init(
		micro.Name("go.micro.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8093"),
		micro.Registry(consulRegistry),
		// 链路追踪绑定
		micro.WrapClient(opentracingFn.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productSrv := pb.NewProductService("go.micro.product.service", srv.Client())

	productAdd := &pb.ProductInfo{
		ProductName:  "王国之泪",
		ProductSku:   "88756",
		ProductPrice: 350,
		ProductDesc:  "NS Game",
		ProductImage: []*pb.ProductImage{
			{
				ImageName: "王国之泪-image",
				ImageCode: "王国之泪-image01",
				ImageUrl:  "xxx",
			},
		},
		ProductSize: []*pb.ProductSize{
			{
				SizeName: "王国之泪-size",
				SizeCode: "王国之泪-size-code",
			},
		},
		ProductSeo: &pb.ProductSeo{
			SeoTitle:    "王国之泪",
			SeoKeywords: "NS",
			SeoDesc:     "good game",
		},
		ProductCategoryId: 1,
	}

	res, err := productSrv.AddProduct(context.Background(), productAdd)

	if err != nil {
		logger.Error(err)
		return
	}

	fmt.Println(res)
}
