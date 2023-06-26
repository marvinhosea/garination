package handlers

import (
	"errors"
	"garination.com/db/internal/core/ports/car"
	"garination.com/db/internal/core/ports/dealership"
	"garination.com/db/internal/core/ports/user"
	"garination.com/db/internal/platform/prom"
	"garination.com/db/sdk/proto"
)

var (
	ErrEmptyRequest = errors.New("empty request")
)

//var _ proto.DatabaseServiceServer = (*Handler)(nil)

type Handler struct {
	userService       user.UserService
	dealershipService dealership.DealershipService
	carService        car.CarService
	promMetrics       prom.Metrics
	proto.UnimplementedDatabaseServiceServer
}

func (h *Handler) mustEmbedUnimplementedDatabaseServiceServer() {
}

func NewHandler(metrics prom.Metrics, userRepo user.UserService, dealershipService dealership.DealershipService, carService car.CarService) *Handler {
	return &Handler{
		carService:        carService,
		userService:       userRepo,
		promMetrics:       metrics,
		dealershipService: dealershipService,
	}
}
