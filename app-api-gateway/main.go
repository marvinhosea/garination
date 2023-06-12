package main

import (
	"fmt"
	"garination.com/gateway/config"
	"garination.com/gateway/internal/platform/cache"
	"garination.com/gateway/internal/platform/casdoor"
	app_db "garination.com/gateway/internal/platform/grpc/app-db"
	"garination.com/gateway/internal/server/http"
	"log"
)

const AsciiArt = `
  __ _ _ __  _ __         __ _  __ _| |_ _____      ____ _ _   _ 
 / _' | '_ \| '_ \ _____ / _' |/ _' | __/ _ \ \ /\ / / _' | | | |
| (_| | |_) | |_) |_____| (_| | (_| | ||  __/\ V  V / (_| | |_| |
 \__,_| .__/| .__/       \__, |\__,_|\__\___| \_/\_/ \__,_|\__, |
      |_|   |_|          |___/                             |___/ 
`

func main() {
	fmt.Print(AsciiArt)

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
