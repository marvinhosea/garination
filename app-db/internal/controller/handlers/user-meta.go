package handlers

import (
	"context"
	"garination.com/db/internal/core/model"
	"garination.com/db/sdk/proto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

var getUserMetaLabel = "GetUserMeta"

func (h *Handler) GetUserMeta(ctx context.Context, req *proto.GetUserMetaRequest) (*proto.GetUserMetaResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getUserMetaLabel))
	h.promMetrics.RequestCount.WithLabelValues(getUserMetaLabel).Inc()
	defer timer.ObserveDuration()

	userMeta, err := h.userService.GetUserMeta(ctx, req.UserId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getUserMetaLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getUserMetaLabel, "success").Inc()
	return &proto.GetUserMetaResponse{
		UserMeta: &proto.UserMetum{
			UserMetaId:   userMeta.UserMetaID,
			UserId:       userMeta.UserID,
			FacebookUrl:  userMeta.FacebookUrl.String,
			TwitterUrl:   userMeta.TwitterUrl.String,
			InstagramUrl: userMeta.InstagramUrl.String,
			LinkedinUrl:  userMeta.LinkedinUrl.String,
			WebsiteUrl:   userMeta.WebsiteUrl.String,
			DealershipId: userMeta.DealershipID.String,
		},
	}, nil
}

var insertUserMetaLabel = "InsertUserMeta"

func (h *Handler) InsertUserMeta(ctx context.Context, req *proto.InsertUserMetaRequest) (*proto.InsertUserMetaResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(insertUserMetaLabel))
	h.promMetrics.RequestCount.WithLabelValues(insertUserMetaLabel).Inc()
	defer timer.ObserveDuration()

	if len(req.UserMeta.UserId) == 0 {
		h.promMetrics.ResponseStatus.WithLabelValues(insertUserMetaLabel, "error").Inc()
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	log.Println(req)

	userMeta := model.UserMetum{
		UserMetaID:   uuid.NewString(),
		UserID:       req.UserMeta.UserId,
		FacebookUrl:  pgtype.Text{String: req.UserMeta.FacebookUrl, Valid: true},
		TwitterUrl:   pgtype.Text{String: req.UserMeta.TwitterUrl, Valid: true},
		InstagramUrl: pgtype.Text{String: req.UserMeta.InstagramUrl, Valid: true},
		LinkedinUrl:  pgtype.Text{String: req.UserMeta.LinkedinUrl, Valid: true},
		WebsiteUrl:   pgtype.Text{String: req.UserMeta.WebsiteUrl, Valid: true},
		DealershipID: pgtype.Text{String: req.UserMeta.DealershipId, Valid: true},
	}

	res, err := h.userService.InsertUserMeta(ctx, userMeta)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(insertUserMetaLabel, "error").Inc()
		log.Println(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(insertUserMetaLabel, "success").Inc()
	return &proto.InsertUserMetaResponse{
		UserMeta: &proto.UserMetum{
			UserMetaId:   res.UserMetaID,
			UserId:       res.UserID,
			FacebookUrl:  res.FacebookUrl.String,
			TwitterUrl:   res.TwitterUrl.String,
			InstagramUrl: res.InstagramUrl.String,
			LinkedinUrl:  res.LinkedinUrl.String,
			WebsiteUrl:   res.WebsiteUrl.String,
			DealershipId: res.DealershipID.String,
		},
	}, err
}

var updateUserMetaLabel = "UpdateUserMeta"

func (h *Handler) UpdateUserMeta(ctx context.Context, req *proto.UpdateUserMetaRequest) (*proto.UpdateUserMetaResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(updateUserMetaLabel))
	h.promMetrics.RequestCount.WithLabelValues(updateUserMetaLabel).Inc()
	defer timer.ObserveDuration()

	userMeta := model.UserMetum{
		FacebookUrl:  pgtype.Text{String: req.UserMeta.FacebookUrl, Valid: true},
		TwitterUrl:   pgtype.Text{String: req.UserMeta.TwitterUrl, Valid: true},
		InstagramUrl: pgtype.Text{String: req.UserMeta.InstagramUrl, Valid: true},
		LinkedinUrl:  pgtype.Text{String: req.UserMeta.LinkedinUrl, Valid: true},
		WebsiteUrl:   pgtype.Text{String: req.UserMeta.WebsiteUrl, Valid: true},
		DealershipID: pgtype.Text{String: req.UserMeta.DealershipId, Valid: true},
		UserID:       req.UserMeta.UserId,
	}

	res, err := h.userService.UpdateUserMeta(ctx, userMeta)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(updateUserMetaLabel, "error").Inc()
		return nil, status.Error(codes.NotFound, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(updateUserMetaLabel, "success").Inc()

	return &proto.UpdateUserMetaResponse{
		UserMeta: &proto.UserMetum{
			UserMetaId:   res.UserMetaID,
			UserId:       res.UserID,
			FacebookUrl:  res.FacebookUrl.String,
			TwitterUrl:   res.TwitterUrl.String,
			InstagramUrl: res.InstagramUrl.String,
			LinkedinUrl:  res.LinkedinUrl.String,
			WebsiteUrl:   res.WebsiteUrl.String,
			DealershipId: res.DealershipID.String,
		},
	}, err
}
