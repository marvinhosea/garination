package dealership

import (
	"context"
	"garination.com/db/internal/platform/postgres"
)

type DealershipPostgresStorage interface {
	GetUserDealership(ctx context.Context, userID string) (postgres.Dealership, error)
	InsertDealership(ctx context.Context, arg postgres.InsertDealershipParams) (postgres.Dealership, error)
	UpdateDealership(ctx context.Context, arg postgres.UpdateDealershipParams) (postgres.Dealership, error)
	GetDealershipByID(ctx context.Context, id string) (postgres.Dealership, error)
	DeleteDealership(ctx context.Context, id string) error
}
