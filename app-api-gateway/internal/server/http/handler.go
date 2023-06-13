package http

import (
	authhttp "garination.com/gateway/internal/core/auth/delivery/http"
	authrepo "garination.com/gateway/internal/core/auth/repository"
	dealershiphttp "garination.com/gateway/internal/core/dealership/delivery/http"
	dealershiprepo "garination.com/gateway/internal/core/dealership/repository"
	dealershipusecase "garination.com/gateway/internal/core/dealership/usecase"
	"garination.com/gateway/internal/core/middleware"
)
import authusecase "garination.com/gateway/internal/core/auth/usecase"

func (s *Server) maphandlers() error {

	// initialise repositories
	authDataBaseRepo := authrepo.NewDatabaseRepo(s.appDbClient)
	authRedisRepo := authrepo.NewRedisRepo(s.redisClient)
	authCasdoorRepo := authrepo.NewCasdoorRepo(s.casdoorClient)
	dealershipRepo := dealershiprepo.NewDealershipDatabaseRepo(s.appDbClient)

	// initialise usecases
	authUsecase := authusecase.NewAuthUsecase(authRedisRepo, authCasdoorRepo, authDataBaseRepo)
	dealershipUsecase := dealershipusecase.NewDealershipUsecase(dealershipRepo)

	// initialise middlewares
	middlewareManager := middleware.NewMiddlewareManager(s.promMetrics)

	// map middlewares
	s.engine.Use(middlewareManager.RequestMetrics)
	s.engine.Use(middlewareManager.Recover)
	s.engine.Use(middlewareManager.Cors)

	v1 := s.engine.Group("/v1")
	auth := v1.Group("/auth")
	dealership := v1.Group("/dealership")

	// initialise handlers
	authHandler := authhttp.NewHandler(authUsecase)
	dealershipHandler := dealershiphttp.NewDealershipHandler(dealershipUsecase)

	// map handlers
	authhttp.MapAuthRoutes(auth, authHandler, middlewareManager)
	dealershiphttp.MapDealershipRoutes(dealership, dealershipHandler, middlewareManager)

	return nil
}
