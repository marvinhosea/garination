package handlers

import (
	"errors"
	"garination.com/db/internal/core/ports/dealership"
	"garination.com/db/internal/core/ports/user"
	"garination.com/db/internal/platform/prom"
	"garination.com/db/sdk/proto"
)

var (
	ErrEmptyRequest = errors.New("empty request")
)

type Handler struct {
	proto.UnimplementedDatabaseServiceServer
	userService       user.UserService
	dealershipService dealership.DealershipService
	promMetrics       prom.Metrics
}

func NewHandler(metrics prom.Metrics, userRepo user.UserService, dealershipService dealership.DealershipService) *Handler {
	return &Handler{
		userService:       userRepo,
		promMetrics:       metrics,
		dealershipService: dealershipService,
	}
}
