package http

import (
	authhttp "garination.com/gateway/internal/core/auth/delivery/http"
	authrepo "garination.com/gateway/internal/core/auth/repository"
	"garination.com/gateway/internal/core/middleware"
)
import authusecase "garination.com/gateway/internal/core/auth/usecase"

func (s *Server) maphandlers() error {

	// initialise repositories
	authDataBaseRepo := authrepo.NewDatabaseRepo(s.appDbClient)
	authRedisRepo := authrepo.NewRedisRepo(s.redisClient)
	authCasdoorRepo := authrepo.NewCasdoorRepo(s.casdoorClient)

	// initialise usecases
	authUsecase := authusecase.NewAuthUsecase(authRedisRepo, authCasdoorRepo, authDataBaseRepo)

	// initialise middlewares
	middlewareManager := middleware.NewMiddlewareManager(s.promMetrics)

	// map middlewares
	s.engine.Use(middlewareManager.RequestMetrics)
	s.engine.Use(middlewareManager.Recover)
	s.engine.Use(middlewareManager.Cors)

	v1 := s.engine.Group("/v1")
	auth := v1.Group("/auth")

	// initialise handlers
	authHandler := authhttp.NewHandler(authUsecase)

	// map handlers
	authhttp.MapAuthRoutes(auth, authHandler, middlewareManager)

	return nil
}
