package cmd

import (
	"garination.com/gateway/config"
	"garination.com/gateway/internal/platform/cache"
	"garination.com/gateway/internal/platform/casdoor"
	app_db "garination.com/gateway/internal/platform/grpc/app-db"
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
	casdoorClient := casdoor.NewCasdoorClient(configurations.Casdoor)

	// start server
	httpServer := http.NewServer(&configurations, redisClient, appDbClient, casdoorClient)
	if err := httpServer.Start(); err != nil {
		log.Panicf("Failed to start server: %v", err)
	}
}
