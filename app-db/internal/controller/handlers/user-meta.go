package handlers

import (
	"context"
	"garination.com/db/internal/core/model"
	"garination.com/db/sdk/proto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Handler) GetUserMeta(ctx context.Context, req *proto.GetUserMetaRequest) (*proto.GetUserMetaResponse, error) {
	userMeta, err := h.userService.GetUserMeta(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

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

func (h *Handler) InsertUserMeta(ctx context.Context, req *proto.InsertUserMetaRequest) (*proto.InsertUserMetaResponse, error) {
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
		return nil, err
	}

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
func (h *Handler) UpdateUserMeta(ctx context.Context, req *proto.UpdateUserMetaRequest) (*proto.UpdateUserMetaResponse, error) {
	userMeta := model.UserMetum{
		UserMetaID:   uuid.NewString(),
		FacebookUrl:  pgtype.Text{String: req.UserMeta.FacebookUrl, Valid: true},
		TwitterUrl:   pgtype.Text{String: req.UserMeta.TwitterUrl, Valid: true},
		InstagramUrl: pgtype.Text{String: req.UserMeta.InstagramUrl, Valid: true},
		LinkedinUrl:  pgtype.Text{String: req.UserMeta.LinkedinUrl, Valid: true},
		WebsiteUrl:   pgtype.Text{String: req.UserMeta.WebsiteUrl, Valid: true},
		DealershipID: pgtype.Text{String: req.UserMeta.DealershipId, Valid: true},
	}

	res, err := h.userService.UpdateUserMeta(ctx, userMeta)
	if err != nil {
		return nil, err
	}

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
