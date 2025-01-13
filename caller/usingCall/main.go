package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

const (
	ServiceName = "stock"
	Endpoint    = "Stock.Quote"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("caller"))
	service.Init()

	// Request message
	payload := make(map[string]interface{})
	payload["symbol"] = "MSFT"

	req := service.Client().NewRequest(
		ServiceName,
		Endpoint,
		payload,
		client.WithContentType("application/json"),
	)

	rsp := &map[string]interface{}{}

	// Call the service
	if err := service.Client().Call(context.Background(), req, rsp); err != nil {
		log.Fatalf("Error calling stock: %v", err)
	}

	out, err := json.Marshal(rsp)
	if err != nil {
		log.Fatalf("Error marshalling response: %v", err)
	}

	fmt.Printf("response:\n %v\n", string(out))

	fmt.Printf("Successfully called %s\n", ServiceName)
}
