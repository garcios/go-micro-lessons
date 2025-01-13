package main

import (
	"github.com/oskiegarcia/go-micro-lessons/mq/handler"
	pb "github.com/oskiegarcia/go-micro-lessons/mq/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("mq"),
		micro.Version("latest"),
	)

	// Register handler
	pb.RegisterMqHandler(srv.Server(), new(handler.Mq))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
