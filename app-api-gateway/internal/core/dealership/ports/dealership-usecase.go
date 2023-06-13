package ports

import (
	"context"
	"garination.com/gateway/internal/core/dealership/dto"
)

type DealerShipUseCase interface {
	GetDealershipByID(ctx context.Context, getDealershipRequest *dto.GetDealershipRequest) (*dto.GetDealershipResponse, error)
	GetDealershipByUserID(ctx context.Context, getDealershipRequest *dto.GetDealershipRequest) (*dto.GetDealershipResponse, error)
	UpdateDealership(ctx context.Context, updateDealershipRequest *dto.UpdateDealershipRequest) (*dto.UpdateDealershipResponse, error)
	CreateDealership(ctx context.Context, createDealershipRequest *dto.CreateDealershipRequest) (*dto.CreateDealershipResponse, error)
	DeleteDealership(ctx context.Context, deleteDealershipRequest *dto.DeleteDealershipRequest) (*dto.DeleteLeadershipResponse, error)
}
