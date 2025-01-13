package main

import (
	"github.com/oskiegarcia/go-micro-lessons/notes/handler"
	pb "github.com/oskiegarcia/go-micro-lessons/notes/proto"
	admin "github.com/oskiegarcia/go-micro-lessons/pkg/service/proto"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

func main() {
	// New Service
	srv := micro.NewService(
		micro.Name("notes"),
		micro.Version("latest"),
	)

	// Initialise service
	srv.Init()

	h := handler.New(srv.Client())
	// Register Handler
	pb.RegisterNotesHandler(srv.Server(), h)
	admin.RegisterAdminHandler(srv.Server(), h)

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
