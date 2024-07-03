package main

import (
	"fmt"

	"user/domain/repository"
	"user/domain/service"
	"user/handler"
	pb "user/proto"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

func main() {
	// 创建服务
	srv := micro.NewService()

	// 服务参数配置以及初始化
	srv.Init(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Address("localhost:8081"),
	)

	// 创建数据库连接
	db, err := gorm.Open("mysql", "root:root123@(127.0.0.1:13306)/market?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}

	db.SingularTable(true)

	// 初始化表，只执行一次
	// rp := repository.NewUserRepository(db)
	// rp.InitTable()

	userDataService := service.NewUserDataService(repository.NewUserRepository(db))

	// Register handler
	if err := pb.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService}); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
