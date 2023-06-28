package cmd

import (
	"garination.com/gateway/config"
	"garination.com/gateway/internal/platform/cache"
	"garination.com/gateway/internal/platform/casdoor"
	app_db "garination.com/gateway/internal/platform/grpc/app-db"
	"garination.com/gateway/internal/platform/prom"
	"garination.com/gateway/internal/platform/wasabi"
	"garination.com/gateway/internal/server/http"
	"log"
)

func Execute() {
	configurations, err := config.LoadConfig()
	if err != nil {
		log.Panicf("Failed to load configurations: %v", err)
	}

	// connect to cache
	redisClient, err := cache.NewRedisClient(configurations.Redis)
	if err != nil {
		log.Panicf("Failed to connect to Redis: %v", err)
	}

	// connect to app-db
	appDbClient, err := app_db.NewDatabaseConn(configurations.AppDb)
	if err != nil {
		log.Panicf("Failed to connect to app-db: %v", err)
	}

	// init casdoor client
	casdoorClient, err := casdoor.NewCasdoorClient(configurations.Casdoor)
	if err != nil {
		log.Panicf("Failed to initialise casdoor client: %v", err)
	}

	// init s3 client
	wasabiClient, err := wasabi.NewWasabiClient(configurations.S3)
	if err != nil {
		log.Panicf("Failed to initialise wasabi client: %v", err)
	}

	deps := http.Dependencies{
		RedisClient:   redisClient,
		AppDbClient:   appDbClient,
		CasdoorClient: casdoorClient,
		WasabiClient:  wasabiClient,
		Config:        &configurations,
		PromMetrics:   prom.NewMetrics(),
	}

	// start server
	httpServer := http.NewServer(deps)
	if err := httpServer.Start(); err != nil {
		log.Panicf("Failed to start server: %v", err)
	}
}
