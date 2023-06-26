package server

import (
	"garination.com/db/config"
	"garination.com/db/internal/controller/handlers"
	"garination.com/db/internal/core/adapters/repository"
	"garination.com/db/internal/core/adapters/service"
	"garination.com/db/internal/platform/postgres"
	"garination.com/db/internal/platform/prom"
	"garination.com/db/sdk/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg config.Config
}

func (s *Server) Run() {

	// prometheus promMetrics
	promMetrics := prom.NewMetrics()

	// create database client
	postgresClient, err := postgres.NewConnection(&s.cfg.Postgres)

	if err != nil {
		log.Panic("error:", err)
	}

	// create  repository
	userRepo := repository.NewUserRepo(postgresClient)
	dealershipRepo := repository.NewDealershipRepo(postgresClient)
	carRepo := repository.NewCarRepository(postgresClient)

	// create user service
	userService := service.NewUserService(userRepo)
	dealershipService := service.NewDealershipService(dealershipRepo, userRepo)
	carService := service.NewCarService(carRepo, userRepo, dealershipRepo)

	// create handler
	grpcHandler := handlers.NewHandler(promMetrics, userService, dealershipService, carService)

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
		log.Println("Starting GRPC server on port", s.cfg.App.Port)
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal("cannot start apps server:", err)
			return
		}
	}()

	// Expose Prometheus promMetrics endpoint
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting Prometheus Metrics EXPOSE server on port", s.cfg.App.MetricsPort)
		if err := http.ListenAndServe(":"+s.cfg.App.MetricsPort, nil); err != nil {
			log.Fatal("Failed to start Prometheus promMetrics server: ", err)
		}
	}()

	<-quit

}

func NewServer(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}
