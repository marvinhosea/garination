package repository

import (
	"context"
	"garination.com/gateway/internal/core/auth/model"
	"garination.com/gateway/internal/core/auth/ports"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
)

type databaseRepo struct {
	client proto.DatabaseServiceClient
}

func (d databaseRepo) GetUserMeta(ctx context.Context, userID string) (*model.UserMetum, error) {
	req := &proto.GetUserMetaRequest{
		UserId: userID,
	}

	res, err := d.client.GetUserMeta(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.UserMetum{
		UserMetaID:   res.UserMeta.UserMetaId,
		UserID:       res.UserMeta.UserId,
		FacebookUrl:  res.UserMeta.FacebookUrl,
		TwitterUrl:   res.UserMeta.TwitterUrl,
		InstagramUrl: res.UserMeta.InstagramUrl,
		LinkedinUrl:  res.UserMeta.LinkedinUrl,
		WebsiteUrl:   res.UserMeta.WebsiteUrl,
		DealershipID: res.UserMeta.DealershipId,
	}, nil
}

func (d databaseRepo) InsertUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	req := &proto.InsertUserMetaRequest{
		UserMeta: &proto.UserMetum{
			UserId:       arg.UserID,
			FacebookUrl:  arg.FacebookUrl,
			TwitterUrl:   arg.TwitterUrl,
			InstagramUrl: arg.InstagramUrl,
			LinkedinUrl:  arg.LinkedinUrl,
			WebsiteUrl:   arg.WebsiteUrl,
			DealershipId: arg.DealershipID,
		},
	}

	res, err := d.client.InsertUserMeta(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.UserMetum{
		UserMetaID:   res.UserMeta.UserMetaId,
		UserID:       res.UserMeta.UserId,
		FacebookUrl:  res.UserMeta.FacebookUrl,
		TwitterUrl:   res.UserMeta.TwitterUrl,
		InstagramUrl: res.UserMeta.InstagramUrl,
		LinkedinUrl:  res.UserMeta.LinkedinUrl,
		WebsiteUrl:   res.UserMeta.WebsiteUrl,
		DealershipID: res.UserMeta.DealershipId,
	}, nil
}

func (d databaseRepo) UpdateUserMeta(ctx context.Context, arg model.UserMetum) (*model.UserMetum, error) {
	req := &proto.UpdateUserMetaRequest{
		UserMeta: &proto.UserMetum{
			UserMetaId:   arg.UserMetaID,
			FacebookUrl:  arg.FacebookUrl,
			TwitterUrl:   arg.TwitterUrl,
			InstagramUrl: arg.InstagramUrl,
			LinkedinUrl:  arg.LinkedinUrl,
			WebsiteUrl:   arg.WebsiteUrl,
			DealershipId: arg.DealershipID,
		},
	}

	res, err := d.client.UpdateUserMeta(ctx, req)
	if err != nil {
		return nil, err
	}

	return &model.UserMetum{
		UserMetaID:   res.UserMeta.UserMetaId,
		UserID:       res.UserMeta.UserId,
		FacebookUrl:  res.UserMeta.FacebookUrl,
		TwitterUrl:   res.UserMeta.TwitterUrl,
		InstagramUrl: res.UserMeta.InstagramUrl,
		LinkedinUrl:  res.UserMeta.LinkedinUrl,
		WebsiteUrl:   res.UserMeta.WebsiteUrl,
		DealershipID: res.UserMeta.DealershipId,
	}, nil
}

func NewDatabaseRepo(client proto.DatabaseServiceClient) ports.DatabaseRepo {
	return &databaseRepo{
		client: client,
	}
}
