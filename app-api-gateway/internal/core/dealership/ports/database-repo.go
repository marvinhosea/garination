package ports

import (
	"context"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
)

type DealershipDatabaseRepo interface {
	GetDealershipByID(ctx context.Context, id string) (*proto.Dealership, error)
	GetDealershipByUserID(ctx context.Context, userID string) (*proto.Dealership, error)
	UpdateDealership(ctx context.Context, dealership *proto.Dealership) (*proto.Dealership, error)
	CreateDealership(ctx context.Context, dealership *proto.Dealership) (*proto.Dealership, error)
	DeleteDealership(ctx context.Context, dealershipId *proto.Dealership) error
}
