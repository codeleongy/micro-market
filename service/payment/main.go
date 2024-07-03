package main

import (
	"micro-market/common"
	"payment/domain/repository"
	"payment/domain/service"
	"payment/handler"
	pb "payment/proto"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/go-micro/plugins/v4/wrapper/monitoring/prometheus"
	opentracingFn "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"

	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	serviceName = "go.micro.service.payment"
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

	// 数据库
	db, err := gorm.Open("mysql", common.GetMysqlURI(common.GetMysqlFromConsul(consulConfig, "mysql")))
	if err != nil {
		logger.Fatal(err)
	}

	db.SingularTable(true)

	defer db.Close()

	repository.NewPaymentRepository(db).InitTable()

	paymentDataService := service.NewPaymentDataService(repository.NewPaymentRepository(db))

	// 开放监控端口
	common.PrometheusBoot(9089)

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address("localhost:8087"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracingFn.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Register handler
	if err := pb.RegisterPaymentHandler(srv.Server(), &handler.Payment{paymentDataService}); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
