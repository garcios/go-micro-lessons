package main

import (
	"context"
	"fmt"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
)

type HelloRequest struct {
}

type HelloResponse struct {
}

func main() {

	// see podman/README.md to see how to start the consul server.
	reg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("hello.client"),
	)
	service.Init()

	// Create a context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"Token": "valid-token", // Normally, this would be a JWT token
	})

	req := client.NewRequest("hello", "Greeter.Hello", &HelloRequest{})
	rsp := &HelloResponse{}

	// Call the service
	if err := service.Client().Call(ctx, req, rsp); err != nil {
		log.Fatalf("Error calling hello service: %v", err)
	}

	fmt.Println("Successfully called hello service")
}
