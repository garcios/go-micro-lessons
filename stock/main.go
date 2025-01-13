package main

import (
	"go-micro.dev/v4/config/source/file"
	"time"

	"github.com/oskiegarcia/go-micro-lessons/stock/handler"
	pb "github.com/oskiegarcia/go-micro-lessons/stock/proto"

	"github.com/patrickmn/go-cache"
	"go-micro.dev/v4"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/logger"
)

type FinageConfig struct {
	Api string `yaml:"api" json:"api"`
	Key string `yaml:"key" json:"key"`
}

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("stock"),
		micro.Version("latest"),
	)

	conf, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("Expected no error but got %v", err)
	}

	// Load file source
	err = conf.Load(file.NewSource(
		file.WithPath("./config/service.json"),
	))

	if err != nil {
		logger.Fatalf("Error loading config: %v", err)
		return
	}

	var appConfig FinageConfig

	err = conf.Get("finage").Scan(&appConfig)
	if err != nil {
		logger.Fatalf("Error scanning config: %v", err)
	}

	if len(appConfig.Api) == 0 {
		logger.Fatal("finage.api config not found")
	}

	if len(appConfig.Key) == 0 {
		logger.Fatal("finage.key config not found")
	}

	// Register handler
	pb.RegisterStockHandler(srv.Server(), &handler.Stock{
		Api:   appConfig.Api,
		Key:   appConfig.Key,
		Cache: cache.New(5*time.Minute, 10*time.Minute),
	})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
