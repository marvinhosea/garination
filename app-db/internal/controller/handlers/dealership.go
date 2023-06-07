package handlers

import (
	"context"
	"garination.com/db/internal/core/model"
	"garination.com/db/sdk/proto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (h *Handler) InsertDealership(ctx context.Context, req *proto.InsertDealershipRequest) (*proto.InsertDealershipResponse, error) {
	dealerShipModel := model.Dealership{
		DealershipID: uuid.NewString(),
		OwnerID:      req.Dealership.OwnerId,
		Name:         req.Dealership.Name,
		DisplayName:  req.Dealership.DisplayName,
		Address:      req.Dealership.Address,
		City:         req.Dealership.City,
		State:        req.Dealership.State,
		Zip:          req.Dealership.Zip,
		Phone:        req.Dealership.Phone,
		Email:        req.Dealership.Email,
		Website:      req.Dealership.Website,
		FacebookUrl:  req.Dealership.FacebookUrl,
		TwitterUrl:   req.Dealership.TwitterUrl,
		InstagramUrl: req.Dealership.InstagramUrl,
		LinkedinUrl:  req.Dealership.LinkedinUrl,
		LogoUrl:      req.Dealership.LogoUrl,
		CoverUrl:     req.Dealership.CoverUrl,
		Description:  req.Dealership.Description,
		CreatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt:    pgtype.Timestamp{},
	}

	res, err := h.userService.InsertDealership(ctx, dealerShipModel)
	if err != nil {
		return nil, err
	}

	return &proto.InsertDealershipResponse{
		Dealership: &proto.Dealership{
			DealershipId: res.DealershipID,
			OwnerId:      res.OwnerID,
			Name:         res.Name,
			DisplayName:  res.DisplayName,
			Address:      res.Address,
			City:         res.City,
			State:        res.State,
			Zip:          res.Zip,
			Phone:        res.Phone,
			Email:        res.Email,
			Website:      res.Website,
			FacebookUrl:  res.FacebookUrl,
			TwitterUrl:   res.TwitterUrl,
			InstagramUrl: res.InstagramUrl,
			LinkedinUrl:  res.LinkedinUrl,
			LogoUrl:      res.LogoUrl,
			CoverUrl:     res.CoverUrl,
			Description:  res.Description,
			CreatedAt:    timestamppb.New(res.CreatedAt.Time),
			UpdatedAt:    timestamppb.New(res.UpdatedAt.Time),
		},
	}, nil
}

func (h *Handler) GetUserDealership(ctx context.Context, req *proto.GetUserDealershipRequest) (*proto.GetUserDealershipResponse, error) {
	userDealership, err := h.userService.GetUserDealership(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserDealershipResponse{
		Dealership: &proto.Dealership{
			DealershipId: userDealership.DealershipID,
			OwnerId:      userDealership.OwnerID,
			Name:         userDealership.Name,
			DisplayName:  userDealership.DisplayName,
			Address:      userDealership.Address,
			City:         userDealership.City,
			State:        userDealership.State,
			Zip:          userDealership.Zip,
			Phone:        userDealership.Phone,
			Email:        userDealership.Email,
			Website:      userDealership.Website,
			FacebookUrl:  userDealership.FacebookUrl,
			TwitterUrl:   userDealership.TwitterUrl,
			InstagramUrl: userDealership.InstagramUrl,
			LinkedinUrl:  userDealership.LinkedinUrl,
			LogoUrl:      userDealership.LogoUrl,
			CoverUrl:     userDealership.CoverUrl,
			Description:  userDealership.Description,
			CreatedAt:    timestamppb.New(userDealership.CreatedAt.Time),
			UpdatedAt:    timestamppb.New(userDealership.UpdatedAt.Time),
		},
	}, err
}
