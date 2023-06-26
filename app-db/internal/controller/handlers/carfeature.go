package handlers

import (
	"context"
	"garination.com/db/internal/platform/postgres"
	"garination.com/db/sdk/proto"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/prometheus/client_golang/prometheus"
)

var insertCarExtraFeatureLabel = "InsertCarExtraFeature"

func (h *Handler) InsertCarExtraFeature(ctx context.Context, request *proto.InsertCarExtraFeatureRequest) (*proto.InsertCarExtraFeatureResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(insertCarExtraFeatureLabel))
	h.promMetrics.RequestCount.WithLabelValues(insertCarExtraFeatureLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.InsertExtraFeatureParams{
		CarID: request.CarExtraFeature.CarId,
		Name:  pgtype.Text{String: request.CarExtraFeature.Name, Valid: true},
		Value: pgtype.Text{String: request.CarExtraFeature.Value, Valid: true},
	}

	feature, err := h.carService.InsertExtraFeature(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(insertCarExtraFeatureLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(insertCarExtraFeatureLabel, "success").Inc()

	return &proto.InsertCarExtraFeatureResponse{
		CarExtraFeatureId: feature.CarExtraFeatureID,
	}, nil
}

var updateCarExtraFeatureLabel = "UpdateCarExtraFeature"

func (h *Handler) UpdateCarExtraFeature(ctx context.Context, request *proto.UpdateCarExtraFeatureRequest) (*proto.UpdateCarExtraFeatureResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(updateCarExtraFeatureLabel))
	h.promMetrics.RequestCount.WithLabelValues(updateCarExtraFeatureLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.UpdateExtraFeatureParams{
		CarExtraFeatureID: request.CarExtraFeature.CarExtraId,
		Name:              pgtype.Text{String: request.CarExtraFeature.Name, Valid: true},
		Value:             pgtype.Text{String: request.CarExtraFeature.Value, Valid: true},
	}

	feature, err := h.carService.UpdateExtraFeature(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(updateCarExtraFeatureLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(updateCarExtraFeatureLabel, "success").Inc()

	return &proto.UpdateCarExtraFeatureResponse{
		CarExtraFeatureId: feature,
	}, nil
}

var deleteCarExtraFeatureLabel = "DeleteCarExtraFeature"

func (h *Handler) DeleteCarExtraFeature(ctx context.Context, request *proto.DeleteCarExtraFeatureRequest) (*proto.DeleteCarExtraFeatureResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(deleteCarExtraFeatureLabel))
	h.promMetrics.RequestCount.WithLabelValues(deleteCarExtraFeatureLabel).Inc()
	defer timer.ObserveDuration()

	res, err := h.carService.DeleteExtraFeature(ctx, request.CarExtraFeatureId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(deleteCarExtraFeatureLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(deleteCarExtraFeatureLabel, "success").Inc()

	return &proto.DeleteCarExtraFeatureResponse{
		CarExtraFeatureId: res.CarExtraFeatureID,
	}, nil
}

var getCarExtraFeatureLabel = "GetCarExtraFeature"

func (h *Handler) GetCarExtraFeaturePaginated(ctx context.Context, request *proto.GetCarExtraFeaturesPaginatedRequest) (*proto.GetCarExtraFeaturesPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarExtraFeatureLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarExtraFeatureLabel).Inc()
	defer timer.ObserveDuration()

	res, err := h.carService.ListExtraFeaturesForCar(ctx, request.CarId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarExtraFeatureLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarExtraFeatureLabel, "success").Inc()

	var features []*proto.CarExtraFeature
	for _, f := range res {
		features = append(features, &proto.CarExtraFeature{
			CarExtraId: f.CarExtraFeatureID,
			Name:       f.Name.String,
			Value:      f.Value.String,
		})
	}

	return &proto.GetCarExtraFeaturesPaginatedResponse{
		CarExtraFeatures: features,
	}, nil
}
