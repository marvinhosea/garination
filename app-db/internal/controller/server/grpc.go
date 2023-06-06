package server

import (
	"garination.com/db/config"
	"garination.com/db/internal/controller/handlers"
	"garination.com/db/internal/core/adapters/repository"
	"garination.com/db/internal/core/adapters/service"
	"garination.com/db/internal/core/adapters/storage"
	"garination.com/db/internal/platform/postgres"
	"garination.com/db/sdk/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg config.Config
}

func (s *Server) Run() {

	// create database client
	postgresClient, err := postgres.NewConnection(&s.cfg.Postgres)

	if err != nil {
		log.Panic("error:", err)
	}

	// Create User storage
	userStorage := storage.NewUserPostgresStorage(postgresClient)

	// create  repository
	userRepo := repository.NewUserRepo(userStorage)

	// create user service
	userService := service.NewUserService(userRepo)

	// create handler
	grpcHandler := handlers.NewHandler(userService)

	// run server
	lis, err := net.Listen("tcp", ":"+s.cfg.App.Port)
	if err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer()
	proto.RegisterDatabaseServiceServer(grpcServer, grpcHandler)

	// Run the server

	// wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal("cannot start apps server:", err)
			return
		}
	}()

	<-quit

}

func NewServer(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}
