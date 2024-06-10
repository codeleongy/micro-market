package main

import (
	"fmt"
	"product/handler"

	"github.com/codeleongy/micro-market/common"
	"github.com/codeleongy/micro-market/service/product/domain/repository"
	"github.com/codeleongy/micro-market/service/product/domain/service"
	pb "github.com/codeleongy/micro-market/service/product/proto"
	"github.com/go-micro/plugins/v4/registry/consul"
	opentracingFn "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	serviceName = "go.micro.product.service"
	version     = "latest"
)

func main() {

	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 18500, "/micro/config")
	if err != nil {
		logger.Error(err)
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"127.0.0.1:18500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.product.service", "localhost:6831")
	if err != nil {
		logger.Error(err)
	}

	defer io.Close()

	opentracing.SetGlobalTracer(t)

	// 数据库设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.User,
		mysqlInfo.Pwd,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.Database,
	))

	if err != nil {
		logger.Error(err)
	}

	defer db.Close()

	db.SingularTable(true)

	// repository.NewProductRepository(db).InitTable()

	productDataService := service.NewProductDataService(repository.NewProductRepository(db))

	// Create service
	srv := micro.NewService()

	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(consulRegistry),
		// 链路追踪绑定
		micro.WrapHandler(opentracingFn.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// Register handler
	if err := pb.RegisterProductHandler(srv.Server(), &handler.Product{productDataService}); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
