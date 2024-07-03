package main

import (
	"webApi/handler"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)

var (
	serviceName = "go.micro.web.gateway"
	version     = "latest"
)

func main() {

	// 注册中心
	consulRegistry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"localhost:18500",
		}
	})

	// Create service
	srv := web.NewService()
	srv.Init(
		web.Name(serviceName),
		web.Version(version),
		web.Address("localhost:8080"),
		web.Registry(consulRegistry),
		web.Handler(handler.GetRouter()),
	)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
