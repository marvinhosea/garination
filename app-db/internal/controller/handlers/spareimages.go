package handlers

import (
	"context"
	"garination.com/db/internal/platform/postgres"
	"garination.com/db/sdk/proto"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var insertSparePartImageRepoLabel = "insertSparePartImageRepo"

func (h *Handler) InsertSparePartImage(ctx context.Context, request *proto.InsertSparePartImageRequest) (*proto.InsertSparePartImageResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(insertCarBrandLabel))
	h.promMetrics.RequestCount.WithLabelValues(insertCarBrandLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.InsertSparePartImageParams{
		SparePartImageID: request.SparePartImageId,
		SparePartID:      request.SparePartId,
		ImageUrl:         pgtype.Text{String: request.ImageUrl, Valid: true},
	}

	image, err := h.spareService.InsertSparePartImage(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(insertSparePartImageRepoLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(insertSparePartImageRepoLabel, "success").Inc()

	return &proto.InsertSparePartImageResponse{
		SparePartImageId: image,
	}, nil

}

var listSparePartImagesBySparePartPagedLabel = "listSparePartImagesBySparePartPaged"

func (h *Handler) ListSparePartImagesBySparePartPaged(ctx context.Context, request *proto.ListSparePartImagesBySparePartPagedRequest) (*proto.ListSparePartImagesBySparePartPagedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(listSparePartImagesBySparePartPagedLabel))
	h.promMetrics.RequestCount.WithLabelValues(listSparePartImagesBySparePartPagedLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListSparePartImagesBySparePartPagedParams{
		SparePartID: request.SparePartId,
		Offset:      request.Offset,
		Limit:       request.Limit,
	}

	images, err := h.spareService.ListSparePartImagesBySparePartPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(listSparePartImagesBySparePartPagedLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(listSparePartImagesBySparePartPagedLabel, "success").Inc()

	var partImages []*proto.SparePartImage = make([]*proto.SparePartImage, len(images), len(images))
	for _, image := range images {
		partImages = append(partImages, &proto.SparePartImage{
			SparePartImageId: image.SparePartImageID,
			SparePartId:      image.SparePartID,
			ImageUrl:         image.ImageUrl.String,
		})
	}

	return &proto.ListSparePartImagesBySparePartPagedResponse{
		SparePartImages: partImages,
	}, nil
}

var listSparePartImagesPagedLabel = "ListSparePartImagesPaged"

func (h *Handler) ListSparePartImagesPaged(ctx context.Context, request *proto.ListSparePartImagesPagedRequest) (*proto.ListSparePartImagesPagedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(listSparePartImagesPagedLabel))
	h.promMetrics.RequestCount.WithLabelValues(listSparePartImagesPagedLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListSparePartImagesPagedParams{
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	images, err := h.spareService.ListSparePartImagesPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(listSparePartImagesPagedLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(listSparePartImagesPagedLabel, "success").Inc()
	var partImages []*proto.SparePartImage = make([]*proto.SparePartImage, len(images), len(images))
	for _, image := range images {
		partImages = append(partImages, &proto.SparePartImage{
			SparePartImageId: image.SparePartImageID,
			SparePartId:      image.SparePartID,
			ImageUrl:         image.ImageUrl.String,
		})
	}

	return &proto.ListSparePartImagesPagedResponse{
		SparePartImages: partImages,
	}, nil
}

var updateSparePartImageLabel = "updateSparePartImage"

func (h *Handler) UpdateSparePartImage(ctx context.Context, request *proto.UpdateSparePartImageRequest) (*proto.UpdateSparePartImageResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(updateSparePartImageLabel))
	h.promMetrics.RequestCount.WithLabelValues(updateSparePartImageLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.UpdateSparePartImageParams{
		SparePartImageID: request.SparePartId,
		SparePartID:      request.SparePartId,
		ImageUrl:         pgtype.Text{String: request.ImageUrl, Valid: true},
	}

	image, err := h.spareService.UpdateSparePartImage(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(updateSparePartImageLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(updateSparePartImageLabel, "success").Inc()

	return &proto.UpdateSparePartImageResponse{
		SparePartImageId: image,
	}, nil
}

var getSparePartImageByIDLabel = "getSparePartImageByID"

func (h *Handler) GetSparePartImageByID(ctx context.Context, request *proto.GetSparePartImageByIDRequest) (*proto.GetSparePartImageByIDResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getSparePartImageByIDLabel))
	h.promMetrics.RequestCount.WithLabelValues(getSparePartImageByIDLabel).Inc()
	defer timer.ObserveDuration()

	image, err := h.spareService.GetSparePartImageById(ctx, request.SparePartImageId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getSparePartImageByIDLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getSparePartImageByIDLabel, "success").Inc()

	return &proto.GetSparePartImageByIDResponse{
		SparePartImage: &proto.SparePartImage{
			SparePartImageId: image.SparePartImageID,
			SparePartId:      image.SparePartImageID,
			ImageUrl:         image.ImageUrl.String,
			CreatedAt:        timestamppb.New(image.CreatedAt.Time),
			UpdatedAt:        timestamppb.New(image.UpdatedAt.Time),
		},
	}, nil
}
