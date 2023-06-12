package handlers

import (
	"context"
	"garination.com/db/internal/platform/postgres"
	"garination.com/db/sdk/proto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var insertDealershipLabel = "InsertDealership"

func (h *Handler) InsertDealership(ctx context.Context, req *proto.InsertDealershipRequest) (*proto.InsertDealershipResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(insertDealershipLabel))
	h.promMetrics.RequestCount.WithLabelValues(insertDealershipLabel).Inc()
	defer timer.ObserveDuration()

	dealerShipModel := postgres.InsertDealershipParams{
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
		UpdatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	res, err := h.dealershipService.InsertDealership(ctx, dealerShipModel)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(insertDealershipLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(insertDealershipLabel, "success").Inc()

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

var getUserDealershipLabel = "GetUserDealership"

func (h *Handler) GetUserDealership(ctx context.Context, req *proto.GetUserDealershipRequest) (*proto.GetDealershipByUserIDResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getUserDealershipLabel))
	h.promMetrics.RequestCount.WithLabelValues(getUserDealershipLabel).Inc()

	defer timer.ObserveDuration()
	userDealership, err := h.dealershipService.GetUserDealership(ctx, req.UserId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getUserDealershipLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getUserDealershipLabel, "success").Inc()
	return &proto.GetDealershipByUserIDResponse{
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

var updateDealershipLabel = "GetDealershipByID"

func (h *Handler) UpdateDealership(ctx context.Context, req *proto.UpdateDealershipRequest) (*proto.UpdateDealershipResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(updateDealershipLabel))
	h.promMetrics.RequestCount.WithLabelValues(updateDealershipLabel).Inc()
	defer timer.ObserveDuration()

	dealership := postgres.UpdateDealershipParams{
		DealershipID: req.Dealership.DealershipId,
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
		UpdatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	res, err := h.dealershipService.UpdateDealership(ctx, dealership)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(updateDealershipLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(updateDealershipLabel, "success").Inc()
	return &proto.UpdateDealershipResponse{
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

var getDealershipByUserIDLabel = "GetDealershipByUserID"

func (h *Handler) GetDealershipByUserId(ctx context.Context, req *proto.GetUserDealershipRequest) (*proto.GetDealershipByUserIDResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getDealershipByUserIDLabel))
	h.promMetrics.RequestCount.WithLabelValues(getDealershipByUserIDLabel).Inc()
	defer timer.ObserveDuration()

	userDealership, err := h.dealershipService.GetUserDealership(ctx, req.UserId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getDealershipByUserIDLabel, "error").Inc()
		return nil, status.Error(codes.NotFound, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getDealershipByUserIDLabel, "success").Inc()
	return &proto.GetDealershipByUserIDResponse{
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
	}, nil
}

var getDealershipByIDLabel = "GetDealershipByID"

func (h *Handler) GetDealershipByID(ctx context.Context, req *proto.GetDealershipByIDRequest) (*proto.GetDealershipByIDResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getDealershipByIDLabel))
	h.promMetrics.RequestCount.WithLabelValues(getDealershipByIDLabel).Inc()
	defer timer.ObserveDuration()

	dealership, err := h.dealershipService.GetDealershipByID(ctx, req.DealershipId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getDealershipByIDLabel, "error").Inc()
		return nil, status.Error(codes.NotFound, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getDealershipByIDLabel, "success").Inc()
	return &proto.GetDealershipByIDResponse{
		Dealership: &proto.Dealership{
			DealershipId: dealership.DealershipID,
			OwnerId:      dealership.OwnerID,
			Name:         dealership.Name,
			DisplayName:  dealership.DisplayName,
			Address:      dealership.Address,
			City:         dealership.City,
			State:        dealership.State,
			Zip:          dealership.Zip,
			Phone:        dealership.Phone,
			Email:        dealership.Email,
			Website:      dealership.Website,
			FacebookUrl:  dealership.FacebookUrl,
			TwitterUrl:   dealership.TwitterUrl,
			InstagramUrl: dealership.InstagramUrl,
			LinkedinUrl:  dealership.LinkedinUrl,
			LogoUrl:      dealership.LogoUrl,
			CoverUrl:     dealership.CoverUrl,
			Description:  dealership.Description,
			CreatedAt:    timestamppb.New(dealership.CreatedAt.Time),
			UpdatedAt:    timestamppb.New(dealership.UpdatedAt.Time),
		},
	}, nil
}

var deleteDealershipLabel = "DeleteDealership"

func (h *Handler) DeleteDealership(ctx context.Context, req *proto.DeleteDealershipRequest) (*proto.DeleteDealershipResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(deleteDealershipLabel))
	h.promMetrics.RequestCount.WithLabelValues(deleteDealershipLabel).Inc()
	defer timer.ObserveDuration()

	err := h.dealershipService.DeleteDealership(ctx, req.DealershipId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(deleteDealershipLabel, "error").Inc()
		return nil, status.Error(codes.NotFound, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(deleteDealershipLabel, "success").Inc()
	return &proto.DeleteDealershipResponse{}, nil
}
