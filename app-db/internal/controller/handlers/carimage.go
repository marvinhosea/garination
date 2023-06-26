package handlers

import (
	"context"
	"garination.com/db/internal/core/dto"
	"garination.com/db/internal/platform/postgres"
	"garination.com/db/sdk/proto"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/prometheus/client_golang/prometheus"
)

var createCarImageLabel = "CreateCarImage"

func (h *Handler) CreateCarImage(ctx context.Context, request *proto.CreateCarImageRequest) (*proto.CreateCarImageResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(createCarImageLabel))
	h.promMetrics.RequestCount.WithLabelValues(createCarImageLabel).Inc()
	defer timer.ObserveDuration()

	req := dto.CarImagesParam{
		CarImages: make([]dto.CarImage, 0),
	}

	for _, image := range request.CarImage {
		req.CarImages = append(req.CarImages, dto.CarImage{
			CarID:    image.CarId,
			ImageUrl: image.Url,
		})
	}

	images, err := h.carService.CreateCarImage(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(createCarImageLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(createCarImageLabel, "success").Inc()

	return &proto.CreateCarImageResponse{
		CarImageIds: images,
	}, nil
}

var updateCarImageLabel = "UpdateCarImage"

func (h *Handler) UpdateCarImage(ctx context.Context, request *proto.UpdateCarImageRequest) (*proto.UpdateCarImageResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(updateCarImageLabel))
	h.promMetrics.RequestCount.WithLabelValues(updateCarImageLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.UpdateCarImageParams{
		CarImageID: request.CarImage.CarImageId,
		ImageUrl: pgtype.Text{
			String: request.CarImage.Url,
			Valid:  true,
		},
	}

	updatedImage, err := h.carService.UpdateCarImage(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(updateCarImageLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(updateCarImageLabel, "success").Inc()

	return &proto.UpdateCarImageResponse{
		CarImageId: updatedImage,
	}, nil
}

var getCarImageLabel = "GetCarImage"

func (h *Handler) GetCarImagePaginated(ctx context.Context, request *proto.GetCarImagesPaginatedRequest) (*proto.GetCarImagesPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarImageLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarImageLabel).Inc()
	defer timer.ObserveDuration()

	images, err := h.carService.ListCarImagesForCar(ctx, request.CarId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarImageLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarImageLabel, "success").Inc()

	var carImages []*proto.CarImage

	for _, image := range images {
		carImages = append(carImages, &proto.CarImage{
			CarImageId: image.CarImageID,
			CarId:      image.CarID,
			Url:        image.ImageUrl.String,
		})
	}

	return &proto.GetCarImagesPaginatedResponse{
		CarImages: carImages,
	}, nil
}

var deleteCarImageLabel = "DeleteCarImage"

func (h *Handler) DeleteCarImage(ctx context.Context, request *proto.DeleteCarImageRequest) (*proto.DeleteCarImageResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(deleteCarImageLabel))
	h.promMetrics.RequestCount.WithLabelValues(deleteCarImageLabel).Inc()
	defer timer.ObserveDuration()

	res, err := h.carService.DeleteCarImage(ctx, request.CarImageId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(deleteCarImageLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(deleteCarImageLabel, "success").Inc()

	return &proto.DeleteCarImageResponse{
		CarImageId: res.CarImageID,
	}, nil
}
