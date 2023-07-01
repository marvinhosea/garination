package handlers

import (
	"context"
	"garination.com/db/internal/platform/postgres"
	"garination.com/db/sdk/proto"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

func sparePostgressToProto(in *postgres.SparePart) (out *proto.SparePart) {
	price, _ := in.Price.Float64Value()

	return &proto.SparePart{
		SparePartId:         in.SparePartID,
		Name:                in.Name.String,
		Description:         in.Description.String,
		Price:               price.Float64,
		Used:                in.Used.Bool,
		CarModel:            in.CarModel.String,
		CarBrand:            in.CarBrand.String,
		OtherCompatibleCars: in.OtherCompatibleCars.Elements,
		CarYear:             in.CarYear.Int32,
		IsUniversal:         in.IsUniversal.Bool,
		Category:            in.Category.String,
		PartNumber:          in.PartNumber.String,
		DealershipId:        in.DealershipID.String,
		DealerId:            in.DealerID.String,
		CreatedAt:           timestamppb.New(in.CreatedAt.Time),
		UpdatedAt:           timestamppb.New(in.UpdatedAt.Time),
		DeletedAt:           timestamppb.New(in.DeletedAt.Time),
	}
}

func sparePostgressArrayToProto(in []postgres.SparePart) (out []*proto.SparePart) {
	for _, i := range in {
		out = append(out, sparePostgressToProto(&i))
	}
	return
}

func spareImagePostgresToProto(in postgres.SparePartImage) (out *proto.SparePartImage) {
	return &proto.SparePartImage{
		SparePartImageId: in.SparePartImageID,
		SparePartId:      in.SparePartID,
		ImageUrl:         in.ImageUrl.String,
	}
}

func spareImagePostgresArrayToProto(in []postgres.SparePartImage) (out []*proto.SparePartImage) {
	for _, i := range in {
		out = append(out, spareImagePostgresToProto(i))
	}
	return
}

var getSparePartByIDLable = "GetSparePartByID"

func (h *Handler) GetSparePartByID(ctx context.Context, request *proto.GetSparePartByIDRequest) (*proto.GetSparePartByIDResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getSparePartByIDLable))
	h.promMetrics.RequestCount.WithLabelValues(getSparePartByIDLable).Inc()
	defer timer.ObserveDuration()

	var err error
	var spare postgres.SparePart

	spare, err = h.spareService.GetSparePartByID(ctx, request.SparePartId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getSparePartByIDLable, "error").Inc()
		return nil, err
	}

	imageReq := postgres.ListSparePartImagesBySparePartPagedParams{
		SparePartID: request.SparePartId,
		Offset:      0,
		Limit:       100,
	}

	images, err := h.spareService.ListSparePartImagesBySparePartPaged(ctx, imageReq)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getSparePartByIDLable, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getSparePartByIDLable, "success").Inc()

	return &proto.GetSparePartByIDResponse{
		SparePart:      sparePostgressToProto(&spare),
		SparePartImage: spareImagePostgresArrayToProto(images),
	}, nil
}

var deleteSparePartLabel = "DeleteSparePart"

func (h *Handler) DeleteSparePart(ctx context.Context, request *proto.DeleteSparePartRequest) (*proto.DeleteSparePartResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(deleteSparePartLabel))
	h.promMetrics.RequestCount.WithLabelValues(deleteSparePartLabel).Inc()
	defer timer.ObserveDuration()

	id, err := h.spareService.DeleteSparePart(ctx, request.SparePartId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(deleteSparePartLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(deleteSparePartLabel, "success").Inc()

	return &proto.DeleteSparePartResponse{
		SparePartId: id,
	}, nil
}

var filterSparePartByBrandPaginatedLabel = "FilterSparePartByBrandPaginated"

func (h *Handler) FilterSparePartByBrandPaginated(ctx context.Context, request *proto.FilterSparePartByBrandPaginatedRequest) (*proto.FilterSparePartByBrandPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(filterSparePartByBrandPaginatedLabel))
	h.promMetrics.RequestCount.WithLabelValues(filterSparePartByBrandPaginatedLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.FilterSparePartsByBrandParams{
		CarBrand: pgtype.Text{String: "",
			Valid: false,
		},
		Limit:  request.Limit,
		Offset: request.Offset,
	}

	spare, err := h.spareService.FilterSparePartsByBrand(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(filterSparePartByBrandPaginatedLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(filterSparePartByBrandPaginatedLabel, "success").Inc()

	return &proto.FilterSparePartByBrandPaginatedResponse{
		SpareParts: sparePostgressArrayToProto(spare),
	}, nil
}

var filterSparePartByCategoryPaginatedLabel = "FilterSparePartByCategoryPaginated"

func (h *Handler) FilterSparePartByCategory(ctx context.Context, request *proto.FilterSparePartByCategoryPaginatedRequest) (*proto.FilterSparePartByCategoryPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(filterSparePartByCategoryPaginatedLabel))
	h.promMetrics.RequestCount.WithLabelValues(filterSparePartByCategoryPaginatedLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.FilterSparePartsByCategoryParams{
		Category: pgtype.Text{String: request.Query, Valid: true},
		Limit:    request.Limit,
		Offset:   request.Offset,
	}

	spares, err := h.spareService.FilterSparePartsByCategory(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(filterSparePartByCategoryPaginatedLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(filterSparePartByCategoryPaginatedLabel, "success").Inc()

	return &proto.FilterSparePartByCategoryPaginatedResponse{
		SpareParts:      sparePostgressArrayToProto(spares),
		Size:            0,
		TotalSpareParts: 0,
	}, nil

}

var filterSparePartByCarModel = "FilterSparePartByCarModel"

func (h *Handler) FilterSparePartByCarModel(ctx context.Context, request *proto.FilterSparePartByCarModelPaginatedRequest) (*proto.FilterSparePartByCarModelPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(filterSparePartByCarModel))
	h.promMetrics.RequestCount.WithLabelValues(filterSparePartByCarModel).Inc()
	defer timer.ObserveDuration()

	req := postgres.FilterSparePartsByModelParams{
		CarModel: pgtype.Text{String: request.Query, Valid: true},
		Limit:    request.Limit,
		Offset:   request.Offset,
	}

	spares, err := h.spareService.FilterSparePartsByModel(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(filterSparePartByCarModel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(filterSparePartByCarModel, "success").Inc()

	return &proto.FilterSparePartByCarModelPaginatedResponse{
		SpareParts:      sparePostgressArrayToProto(spares),
		Size:            0,
		TotalSpareParts: 0,
	}, nil
}

var insertSparePartLabel = "InsertSparePart"

func (h *Handler) InsertSparePart(ctx context.Context, request *proto.InsertSparePartRequest) (*proto.InsertSparePartResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(insertSparePartLabel))
	h.promMetrics.RequestCount.WithLabelValues(insertSparePartLabel).Inc()
	defer timer.ObserveDuration()

	sparePrice := pgtype.Numeric{}
	err := sparePrice.UnmarshalJSON([]byte(strconv.FormatFloat(request.Price, 'f', 2, 64)))
	if err != nil {
		return nil, err
	}

	req := postgres.InsertSparePartParams{
		Name:                pgtype.Text{String: request.Name, Valid: true},
		Description:         pgtype.Text{String: request.Description, Valid: true},
		Price:               sparePrice,
		Used:                pgtype.Bool{Bool: request.Used, Valid: true},
		CarModel:            pgtype.Text{String: request.CarModel, Valid: true},
		CarBrand:            pgtype.Text{String: request.CarBrand, Valid: true},
		OtherCompatibleCars: pgtype.Array[string]{Elements: request.OtherCompatibleCars, Valid: true},
		CarYear:             pgtype.Int4{Int32: int32(request.CarYear), Valid: true},
		IsUniversal:         pgtype.Bool{Bool: request.IsUniversal, Valid: true},
		Category:            pgtype.Text{String: request.Category, Valid: true},
		PartNumber:          pgtype.Text{String: request.PartNumber, Valid: true},
		DealershipID:        pgtype.Text{String: request.DealershipId, Valid: true},
		DealerID:            pgtype.Text{String: request.DealerId, Valid: true},
	}

	id, err := h.spareService.InsertSparePart(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(insertSparePartLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(insertSparePartLabel, "success").Inc()

	return &proto.InsertSparePartResponse{
		SparePartId: id,
	}, nil
}

var listSparePartsByDealerPagedLabel = "ListSparePartsByDealerPaged"

func (h *Handler) ListSparePartsByDealerPaged(ctx context.Context, request *proto.ListSparePartsByDealerPagedRequest) (*proto.ListSparePartsByDealerPagedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(listSparePartsByDealerPagedLabel))
	h.promMetrics.RequestCount.WithLabelValues(listSparePartsByDealerPagedLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListSparePartsByDealerPagedParams{
		DealerID: pgtype.Text{String: request.DealerId, Valid: true},
		Limit:    request.Limit,
		Offset:   request.Offset,
	}

	spares, err := h.spareService.ListSparePartsByDealerPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(listSparePartsByDealerPagedLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(listSparePartsByDealerPagedLabel, "success").Inc()

	return &proto.ListSparePartsByDealerPagedResponse{
		SpareParts:      sparePostgressArrayToProto(spares),
		Size:            0,
		TotalSpareParts: 0,
	}, nil
}

var listSparePartsByDealershipPagedLabel = "ListSparePartsByDealershipPaged"

func (h *Handler) ListSparePartsByDealershipPaged(ctx context.Context, request *proto.ListSparePartsByDealershipPagedRequest) (*proto.ListSparePartsByDealershipPagedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(listSparePartsByDealershipPagedLabel))
	h.promMetrics.RequestCount.WithLabelValues(listSparePartsByDealershipPagedLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListSparePartsByDealershipPagedParams{
		DealershipID: pgtype.Text{String: request.DealershipId, Valid: true},
		Limit:        request.Limit,
		Offset:       request.Offset,
	}

	spares, err := h.spareService.ListSparePartsByDealershipPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(listSparePartsByDealershipPagedLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(listSparePartsByDealershipPagedLabel, "success").Inc()

	return &proto.ListSparePartsByDealershipPagedResponse{
		SpareParts:      sparePostgressArrayToProto(spares),
		Size:            0,
		TotalSpareParts: 0,
	}, nil
}

var listSparePartsPaged = "ListSparePartsPaged"

func (h *Handler) ListSparePartsPaged(ctx context.Context, request *proto.ListSparePartsPagedRequest) (*proto.ListSparePartsPagedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(listSparePartsPaged))
	h.promMetrics.RequestCount.WithLabelValues(listSparePartsPaged).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListSparePartsPagedParams{
		Limit:  request.Limit,
		Offset: request.Offset,
	}

	spares, err := h.spareService.ListSparePartsPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(listSparePartsPaged, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(listSparePartsPaged, "success").Inc()

	return &proto.ListSparePartsPagedResponse{
		SpareParts:      sparePostgressArrayToProto(spares),
		Size:            0,
		TotalSpareParts: 0,
	}, nil
}

var searchSparePartsPagedLabel = "SearchSparePartsPaged"

func (h *Handler) SearchSparePartsPaged(ctx context.Context, request *proto.SearchSparePartsPagedRequest) (*proto.SearchSparePartsPagedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(searchSparePartsPagedLabel))
	h.promMetrics.RequestCount.WithLabelValues(searchSparePartsPagedLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.SearchSparePartsPagedParams{
		Lower:  request.Query,
		Limit:  request.Limit,
		Offset: request.Offset,
	}

	spares, err := h.spareService.SearchSparePartsPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(searchSparePartsPagedLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(searchSparePartsPagedLabel, "success").Inc()

	return &proto.SearchSparePartsPagedResponse{
		SpareParts:      sparePostgressArrayToProto(spares),
		Size:            0,
		TotalSpareParts: 0,
	}, nil
}

var updateSparePartLabel = "UpdateSparePart"

func (h *Handler) UpdateSparePart(ctx context.Context, request *proto.UpdateSparePartRequest) (*proto.UpdateSparePartResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(updateSparePartLabel))
	h.promMetrics.RequestCount.WithLabelValues(updateSparePartLabel).Inc()
	defer timer.ObserveDuration()

	sparePrice := pgtype.Numeric{}
	err := sparePrice.UnmarshalJSON([]byte(string(request.Price)))
	if err != nil {
		return nil, err
	}

	req := postgres.UpdateSparePartParams{
		SparePartID:         request.SparePartId,
		Name:                pgtype.Text{String: request.Name, Valid: true},
		Description:         pgtype.Text{String: request.Description, Valid: true},
		Price:               sparePrice,
		Used:                pgtype.Bool{Bool: true, Valid: true},
		CarModel:            pgtype.Text{String: request.CarModel, Valid: true},
		CarBrand:            pgtype.Text{String: request.CarBrand, Valid: true},
		OtherCompatibleCars: pgtype.Array[string]{Elements: request.OtherCompatibleCars, Valid: true},
		CarYear:             pgtype.Int4{Int32: request.CarYear, Valid: true},
		IsUniversal:         pgtype.Bool{Bool: request.IsUniversal, Valid: true},
		Category:            pgtype.Text{String: request.Category, Valid: true},
		PartNumber:          pgtype.Text{String: request.PartNumber, Valid: true},
	}

	updatedId, err := h.spareService.UpdateSparePart(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(updateSparePartLabel, "error").Inc()
		return nil, err
	}

	h.promMetrics.ResponseStatus.WithLabelValues(updateSparePartLabel, "success").Inc()

	return &proto.UpdateSparePartResponse{
		SparePartId: updatedId,
	}, nil
}
