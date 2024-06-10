package main

import (

	"webApi/handler"
	pb "webApi/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

)

var (
	service = "webapi"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterWebApiHandler(srv.Server(), new(handler.WebApi)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
