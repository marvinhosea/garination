package handlers

import (
	"errors"
	"garination.com/db/internal/core/ports"
	"garination.com/db/sdk/proto"
)

var (
	ErrEmptyRequest = errors.New("empty request")
)

type Handler struct {
	proto.UnimplementedDatabaseServiceServer
	userService ports.UserService
}

func NewHandler(userRepo ports.UserService) *Handler {
	return &Handler{
		userService: userRepo,
	}
}
