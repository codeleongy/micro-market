package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/codeleongy/micro-market/api/cartApi/handler"
	pb "github.com/codeleongy/micro-market/api/cartApi/proto"
	"github.com/codeleongy/micro-market/common"
	cart "github.com/codeleongy/micro-market/service/cart/proto"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/wrapper/select/roundrobin"
	opentracingFn "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	serviceName = "go.micro.service.cartapi"
	version     = "latest"
)

func main() {

	// 注册中心
	consulRegistry := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{"localhost:18500"}
	})

	// 链路追踪
	t, io, err := common.NewTracer(serviceName, "localhost:6831")
	if err != nil {
		logger.Error(err)
		return
	}

	defer io.Close()

	opentracing.SetGlobalTracer(t)

	// 熔断器
	hystrixHandler := hystrix.NewStreamHandler()
	// 启动熔断器
	hystrixHandler.Start()
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9096"), hystrixHandler)
		if err != nil {
			logger.Error(err)
		}
	}()

	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address("0.0.0.0:8085"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracingFn.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracingFn.NewClientWrapper(opentracing.GlobalTracer())),
		// 添加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		// 添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	cartService := cart.NewCartService("go.micro.service.cart", srv.Client())

	// Register handler
	if err := pb.RegisterCartApiHandler(srv.Server(), &handler.CartApi{CartService: cartService}); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		// 正常执行的逻辑
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
