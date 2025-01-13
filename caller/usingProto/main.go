package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/oskiegarcia/go-micro-lessons/stock/proto"
	"go-micro.dev/v4"
	"log"
)

const (
	ServiceName = "stock"
)

func main() {
	// Create a new service
	service := micro.NewService(micro.Name("caller"))
	service.Init()

	req := pb.QuoteRequest{
		Symbol: "GOOGL",
	}

	stockService := pb.NewStockService(ServiceName, service.Client())

	rsp, err := stockService.Quote(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.Marshal(rsp)
	if err != nil {
		log.Fatalf("Error marshalling response: %v", err)
	}

	fmt.Printf("response:\n %v\n", string(out))

	fmt.Printf("Successfully called %s\n", ServiceName)

}
