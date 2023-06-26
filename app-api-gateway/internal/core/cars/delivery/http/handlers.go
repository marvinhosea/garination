package http

import (
	"garination.com/gateway/internal/core/cars/ports"
)

type handlerImpl struct {
	carUsecase ports.CarUsercase
}

func NewHandler(carUsecase ports.CarUsercase) ports.Handler {
	return &handlerImpl{
		carUsecase: carUsecase,
	}
}
