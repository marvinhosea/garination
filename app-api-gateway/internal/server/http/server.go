package http

import (
	"garination.com/gateway/config"
	"garination.com/gateway/internal/platform/casdoor"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"garination.com/gateway/internal/platform/prom"
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

// Server struct
type Server struct {
	engine        *gin.Engine
	cfg           *config.Config
	redisClient   *redis.Client
	appDbClient   proto.DatabaseServiceClient
	promMetrics   prom.Metrics
	casdoorClient casdoor.CasdoorClient
}

// NewServer returns a new server instance
func NewServer(cfg *config.Config, redisClient *redis.Client, appDbClient proto.DatabaseServiceClient, casdoorClient casdoor.CasdoorClient) *Server {
	engine := gin.Default()
	return &Server{
		engine:        engine,
		cfg:           cfg,
		redisClient:   redisClient,
		appDbClient:   appDbClient,
		promMetrics:   prom.NewMetrics(),
		casdoorClient: casdoorClient,
	}
}

// Start starts the server
func (s *Server) serve() error {
	server := &http.Server{
		Addr:           ":" + s.cfg.App.Port,
		ReadTimeout:    time.Minute * 5,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: maxHeaderBytes,
	}

	// Expose Prometheus promMetrics endpoint
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting Prometheus Metrics EXPOSE server on port", s.cfg.App.MetricsPort)
		if err := http.ListenAndServe(":"+s.cfg.App.MetricsPort, nil); err != nil {
			log.Fatal("Failed to start Prometheus promMetrics server: ", err)
		}
	}()

	go func() {
		log.Printf("Server is listening on PORT: %s\n", s.cfg.App.Port)
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
