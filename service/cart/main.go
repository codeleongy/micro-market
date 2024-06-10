package main

import (
	"cart/domain/repository"
	"cart/domain/service"
	"fmt"

	"cart/handler"
	pb "cart/proto"

	"github.com/codeleongy/micro-market/common"
	"github.com/go-micro/plugins/v4/registry/consul"
	ratelimit "github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber"
	opentracingFn "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	serviceName = "go.micro.service.cart"
	version     = "latest"
	QPS         = 100
)

func main() {

	// 配置中心
	consulConfig, err := common.GetConsulConfig("localhost", 18500, "/micro/config")
	if err != nil {
		logger.Error(err)
		return
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"localhost:18500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")

	if err != nil {
		logger.Error(err)
		return
	}

	defer io.Close()

	opentracing.SetGlobalTracer(t)

	// 数据库初始化
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
		return
	}

	defer db.Close()

	db.SingularTable(true)

	// 只执行一次
	// repository.NewCartRepository(db).InitTable()

	cartDataService := service.NewCartDataService(repository.NewCartRepository(db))

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address("localhost:8084"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracingFn.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 添加限流器
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	// Register handler
	if err := pb.RegisterCartHandler(srv.Server(), &handler.Cart{cartDataService}); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
