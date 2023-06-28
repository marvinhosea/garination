package handlers

import (
	"context"
	"garination.com/db/internal/platform/postgres"
	"garination.com/db/sdk/proto"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/prometheus/client_golang/prometheus"
)

var insertCarBrandLabel = "InsertCarBrand"

func (h *Handler) InsertCarBrand(ctx context.Context, request *proto.InsertCarBrandRequest) (*proto.InsertCarBrandResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(insertCarBrandLabel))
	h.promMetrics.RequestCount.WithLabelValues(insertCarBrandLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.CreateCarBrandParams{
		Name:    pgtype.Text{String: request.CarBrand.Name, Valid: true},
		LogoUrl: pgtype.Text{String: request.CarBrand.LogoUrl, Valid: true},
		Country: pgtype.Text{String: request.CarBrand.Country, Valid: true},
	}

	brand, err := h.carService.CreateCarBrand(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(insertCarBrandLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(insertCarBrandLabel, "success").Inc()

	return &proto.InsertCarBrandResponse{
		CarBrandId: brand.BrandID,
	}, nil

}

var updateCarBrandLabel = "UpdateCarBrand"

func (h *Handler) UpdateCarBrand(ctx context.Context, request *proto.UpdateCarBrandRequest) (*proto.UpdateCarBrandResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(updateCarBrandLabel))
	h.promMetrics.RequestCount.WithLabelValues(updateCarBrandLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.UpdateCarBrandParams{
		BrandID: request.CarBrand.BrandId,
		Name:    pgtype.Text{String: request.CarBrand.Name, Valid: true},
		LogoUrl: pgtype.Text{String: request.CarBrand.LogoUrl, Valid: true},
	}

	brand, err := h.carService.UpdateCarBrand(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(updateCarBrandLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(updateCarBrandLabel, "success").Inc()

	return &proto.UpdateCarBrandResponse{
		CarBrandId: brand.BrandID,
	}, nil
}

var getCarBrandByIDLabel = "GetCarBrandByID"

func (h *Handler) GetCarBrandByID(ctx context.Context, request *proto.GetCarBrandByIDRequest) (*proto.GetCarBrandByIDResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarBrandByIDLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarBrandByIDLabel).Inc()
	defer timer.ObserveDuration()

	brand, err := h.carService.GetCarBrandById(ctx, request.CarBrandId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarBrandByIDLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarBrandByIDLabel, "success").Inc()

	return &proto.GetCarBrandByIDResponse{
		CarBrand: &proto.CarBrand{
			BrandId: brand.BrandID,
			Name:    brand.Name.String,
			LogoUrl: brand.LogoUrl.String,
		},
	}, nil
}

var getCarBrandsPaginatedLabel = "GetCarBrandsPaginated"

func (h *Handler) GetCarBrandsPaginated(ctx context.Context, request *proto.GetCarBrandsPaginatedRequest) (*proto.GetCarBrandsPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarBrandsPaginatedLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarBrandsPaginatedLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListCarBrandsPagedParams{
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	brands, err := h.carService.ListCarBrandsPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarBrandsPaginatedLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarBrandsPaginatedLabel, "success").Inc()

	var carBrands []*proto.CarBrand
	for _, brand := range brands {
		carBrands = append(carBrands, &proto.CarBrand{
			BrandId: brand.BrandID,
			Name:    brand.Name.String,
			LogoUrl: brand.LogoUrl.String,
		})
	}

	return &proto.GetCarBrandsPaginatedResponse{
		CarBrands: carBrands,
	}, nil
}

var deleteCarBrandLabel = "DeleteCarBrand"

func (h *Handler) DeleteCarBrand(ctx context.Context, request *proto.DeleteCarBrandRequest) (*proto.DeleteCarBrandResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(deleteCarBrandLabel))
	h.promMetrics.RequestCount.WithLabelValues(deleteCarBrandLabel).Inc()
	defer timer.ObserveDuration()

	res, err := h.carService.DeleteCarBrand(ctx, request.CarBrandId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(deleteCarBrandLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(deleteCarBrandLabel, "success").Inc()

	return &proto.DeleteCarBrandResponse{
		CarBrandId: res.BrandID,
	}, nil
}
