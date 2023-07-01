package http

import (
	"garination.com/gateway/config"
	"garination.com/gateway/internal/platform/casdoor"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"garination.com/gateway/internal/platform/prom"
	"garination.com/gateway/internal/platform/wasabi"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	readTimeout    = 5
	writeTimeout   = 10
)

type Dependencies struct {
	RedisClient   *redis.Client
	AppDbClient   proto.DatabaseServiceClient
	CasdoorClient casdoor.CasdoorClient
	WasabiClient  *wasabi.Client
	Config        *config.Config
	PromMetrics   prom.Metrics
}

// Server struct
type Server struct {
	engine       *gin.Engine
	dependencies *Dependencies
}

// NewServer returns a new server instance
func NewServer(dependencies Dependencies) *Server {
	engine := gin.Default()
	return &Server{
		engine:       engine,
		dependencies: &dependencies,
	}
}

// Start starts the server
func (s *Server) serve() error {
	server := &http.Server{
		Addr:           ":" + s.dependencies.Config.App.Port,
		ReadTimeout:    time.Minute * 5,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: maxHeaderBytes,
	}

	// Expose Prometheus promMetrics endpoint
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting Prometheus Metrics EXPOSE server on port", s.dependencies.Config.App.MetricsPort)
		if err := http.ListenAndServe(":"+s.dependencies.Config.App.MetricsPort, nil); err != nil {
			log.Fatal("Failed to start Prometheus promMetrics server: ", err)
		}
	}()

	go func() {
		log.Printf("Server is listening on PORT: %s\n", s.dependencies.Config.App.Port)
		if err := s.engine.Run(server.Addr); err != nil {
			log.Fatalf("Error starting Server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	// wait for interrupt signal to gracefully shutdown the server with
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server...")
	return nil
}

// Start starts the server
func (s *Server) Start() error {
	// map handlers
	if err := s.maphandlers(); err != nil {
		return err
	}
	return s.serve()
}
