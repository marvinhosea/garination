package cmd

import (
	"garination.com/db/config"
	"garination.com/db/internal/controller/server"
	"log"
)

func Execute() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Panic("config error:", err)
	}

	grpcServer := server.NewServer(cfg)
	grpcServer.Run()
}
