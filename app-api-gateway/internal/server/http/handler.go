package http

import (
	authhttp "garination.com/gateway/internal/core/auth/delivery/http"
	authrepo "garination.com/gateway/internal/core/auth/repository"
	carhttp "garination.com/gateway/internal/core/cars/delivery/http"
	carrepo "garination.com/gateway/internal/core/cars/repository"
	carusecase "garination.com/gateway/internal/core/cars/usecase"
	"garination.com/gateway/internal/core/common"
	dealershiphttp "garination.com/gateway/internal/core/dealership/delivery/http"
	dealershiprepo "garination.com/gateway/internal/core/dealership/repository"
	dealershipusecase "garination.com/gateway/internal/core/dealership/usecase"
	"garination.com/gateway/internal/core/middleware"
)
import authusecase "garination.com/gateway/internal/core/auth/usecase"

func (s *Server) maphandlers() error {

	// initialise repositories
	authDataBaseRepo := authrepo.NewDatabaseRepo(s.dependencies.AppDbClient)
	authRedisRepo := authrepo.NewRedisRepo(s.dependencies.RedisClient)
	authCasdoorRepo := authrepo.NewCasdoorRepo(s.dependencies.CasdoorClient)
	dealershipRepo := dealershiprepo.NewDealershipDatabaseRepo(s.dependencies.AppDbClient)
	carRepo := carrepo.NewCarRepo(s.dependencies.AppDbClient)
	wasabiRepo := common.NewWasabiRepo(s.dependencies.WasabiClient)

	// initialise usecases
	authUsecase := authusecase.NewAuthUsecase(authRedisRepo, authCasdoorRepo, authDataBaseRepo)
	dealershipUsecase := dealershipusecase.NewDealershipUsecase(dealershipRepo)
	carUsecase := carusecase.NewUsecase(carRepo, wasabiRepo)

	// initialise middlewares
	middlewareManager := middleware.NewMiddlewareManager(s.dependencies.PromMetrics, s.dependencies.CasdoorClient)

	// map middlewares
	s.engine.Use(middlewareManager.RequestMetrics)
	s.engine.Use(middlewareManager.Recover)
	s.engine.Use(middlewareManager.Cors)

	v1 := s.engine.Group("/v1")
	auth := v1.Group("/auth")
	dealership := v1.Group("/dealership")
	cars := v1.Group("/cars")

	// initialise handlers
	authHandler := authhttp.NewHandler(authUsecase)
	dealershipHandler := dealershiphttp.NewDealershipHandler(dealershipUsecase)
	carHandler := carhttp.NewHandler(carUsecase)

	// map handlers
	authhttp.MapAuthRoutes(auth, authHandler, middlewareManager)
	dealershiphttp.MapDealershipRoutes(dealership, dealershipHandler, middlewareManager)
	carhttp.MapCarRoutes(cars, carHandler, middlewareManager)

	return nil
}
