package handlers

import (
	"context"
	"garination.com/db/internal/core/model"
	"garination.com/db/sdk/proto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) InsertDealership(ctx context.Context, req *proto.InsertDealershipRequest) (*proto.InsertDealershipResponse, error) {
	dealerShipModel := model.Dealership{
		DealershipID:  uuid.NewString(),
		OwnerID:       req.Dealership.OwnerId,
		Name:          req.Dealership.Name,
		DisplayName:   req.Dealership.DisplayName,
		Address:       req.Dealership.Address,
		City:          req.Dealership.City,
		State:         req.Dealership.State,
		Zip:           req.Dealership.Zip,
		Phone:         req.Dealership.Phone,
		Email:         req.Dealership.Email,
		Website:       req.Dealership.Website,
		FacebookUrl:   req.Dealership.FacebookUrl,
		TwitterUrl:    req.Dealership.TwitterUrl,
		InstagramUrl:  req.Dealership.InstagramUrl,
		LinkedinUrl:   req.Dealership.LinkedinUrl,
		LogoUrl:       req.Dealership.LogoUrl,
		CoverUrl:      req.Dealership.CoverUrl,
		Description:   req.Dealership.Description,
		CreatedAt:     pgtype.,
		UpdatedAt:     pgtype.Timestamp{},
	}
}
