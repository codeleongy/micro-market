package main

import (
	"fmt"

	"category/domain/repository"
	"category/domain/service"
	"category/handler"
	pb "category/proto"
	"micro-market/common"

	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	serviceName = "go.micro.service.category"
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

	// Create service
	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
		// 设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 添加consul为注册中心
		micro.Registry(consulRegistry),
	)

	// 获取mysql配置，路径中不用带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.User,
		mysqlInfo.Pwd,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.Database,
	))
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// 禁止复表
	db.SingularTable(true)

	// 初始化表
	// rp := repository.NewCategoryRepository(db)
	// rp.InitTable()

	// 初始化服务
	srv.Init()

	categoryDataService := service.NewCategoryDataService(repository.NewCategoryRepository(db))

	// Register handler
	if err := pb.RegisterCategoryHandler(srv.Server(), &handler.Category{categoryDataService}); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
