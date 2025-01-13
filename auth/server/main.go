package main

import (
	"context"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

type Greeter struct{}

type HelloRequest struct {
}

type HelloResponse struct {
}

func (g *Greeter) Hello(ctx context.Context, req *HelloRequest, rsp *HelloResponse) error {
	fmt.Println("Hello service was called")
	// Business logic goes here...
	return nil
}

func main() {

	// see podman/README.md to see how to start the consul server.
	reg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("hello"),
		micro.WrapHandler(AuthMiddleware()),
	)

	service.Init()

	if err := micro.RegisterHandler(service.Server(), new(Greeter)); err != nil {
		fmt.Println(err)
		return
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
