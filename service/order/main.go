package main

import (
	"micro-market/common"
	"order/domain/repository"
	"order/domain/service"
	"order/handler"
	pb "order/proto"

	"github.com/go-micro/plugins/v4/wrapper/monitoring/prometheus"

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
	serviceName = "go.micro.service.order"
	version     = "latest"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("localhost", 18500, "/micro/config")
	if err != nil {
		logger.Fatal(err)
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"localhost:18500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer(serviceName, "localhost:6831")
	if err != nil {
		logger.Fatal(err)
	}

	defer io.Close()

	opentracing.SetGlobalTracer(t)

	// 暴露监控地址
	common.PrometheusBoot(9092)

	// 数据库初始化
	db, err := gorm.Open("mysql", common.GetMysqlURI(common.GetMysqlFromConsul(consulConfig, "mysql")))
	if err != nil {
		logger.Fatal(err)
	}
	db.SingularTable(true)

	// 只执行一次，用于创建表
	repository.NewOrderRepository(db).InitTable()

	orderDataService := service.NewOrderDataService(repository.NewOrderRepository(db))

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address("localhost:8086"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracingFn.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Register handler
	if err := pb.RegisterOrderHandler(srv.Server(), &handler.Order{OrderDataService: orderDataService}); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
