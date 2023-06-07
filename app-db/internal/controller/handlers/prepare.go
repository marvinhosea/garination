package handlers

import (
	"errors"
	"garination.com/db/internal/core/ports"
	"garination.com/db/internal/platform/prom"
	"garination.com/db/sdk/proto"
)

var (
	ErrEmptyRequest = errors.New("empty request")
)

type Handler struct {
	proto.UnimplementedDatabaseServiceServer
	userService ports.UserService
	promMetrics prom.Metrics
}

func NewHandler(metrics prom.Metrics, userRepo ports.UserService) *Handler {
	return &Handler{
		userService: userRepo,
		promMetrics: metrics,
	}
}
