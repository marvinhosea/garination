package handlers

import (
	"context"
	"garination.com/db/internal/core/dto"
	"garination.com/db/internal/platform/postgres"
	"garination.com/db/sdk/proto"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
)

var healthCheckLabel = "HealthCheck"

func (h *Handler) HealthCheck(context.Context, *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(healthCheckLabel))
	h.promMetrics.RequestCount.WithLabelValues(healthCheckLabel).Inc()
	defer timer.ObserveDuration()

	// keep metrics
	h.promMetrics.ResponseStatus.WithLabelValues(healthCheckLabel, "success").Inc()

	return &proto.HealthCheckResponse{
		Status: "OK",
	}, nil
}

var insertCarLabel = "InsertCar"

func (h *Handler) InsertCar(ctx context.Context, request *proto.InsertCarRequest) (*proto.InsertCarResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(insertCarLabel))
	h.promMetrics.RequestCount.WithLabelValues(insertCarLabel).Inc()
	defer timer.ObserveDuration()

	carPrice := &pgtype.Numeric{}
	err := carPrice.UnmarshalJSON([]byte(strconv.FormatFloat(request.Car.Price, 'f', 2, 64)))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := postgres.InsertCarParams{
		BrandID:                   request.Car.BrandId,
		Model:                     pgtype.Text{String: request.Car.Model, Valid: true},
		Year:                      pgtype.Int4{Int32: request.Car.Year, Valid: true},
		Price:                     *carPrice,
		Mileage:                   pgtype.Int4{Int32: request.Car.Mileage, Valid: true},
		Color:                     pgtype.Text{String: request.Car.Color, Valid: true},
		Transmission:              pgtype.Text{String: request.Car.Transmission, Valid: true},
		FuelType:                  pgtype.Text{String: request.Car.FuelType, Valid: true},
		EngineCapacity:            pgtype.Int4{Int32: request.Car.EngineCapacity, Valid: true},
		Description:               pgtype.Text{String: request.Car.Description, Valid: true},
		DealershipID:              pgtype.Text{String: request.Car.DealershipId, Valid: true},
		DealerID:                  pgtype.Text{String: request.Car.DealerId, Valid: true},
		ComfortSpecifications:     pgtype.Array[string]{Elements: request.Car.PerformanceSpecifications, Valid: true},
		SafetySpecifications:      pgtype.Array[string]{Elements: request.Car.SafetySpecifications, Valid: true},
		PerformanceSpecifications: pgtype.Array[string]{Elements: request.Car.PerformanceSpecifications, Valid: true},
		HorsePower:                pgtype.Int4{Int32: request.Car.HorsePower, Valid: true},
		Torque:                    pgtype.Int4{Int32: request.Car.Torque, Valid: true},
		TorqueRpm:                 pgtype.Int4{Int32: request.Car.TorqueRpm, Valid: true},
		Ownership:                 pgtype.Text{String: request.Car.Ownership, Valid: true},
		Title:                     pgtype.Text{String: request.Car.Title, Valid: true},
	}

	insertCar, err := h.carService.InsertCar(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(insertCarLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(insertCarLabel, "success").Inc()

	return &proto.InsertCarResponse{
		CarId: insertCar.CarID,
	}, nil

}

var updateCarLabel = "UpdateCar"

func (h *Handler) UpdateCar(ctx context.Context, request *proto.UpdateCarRequest) (*proto.UpdateCarResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(updateCarLabel))
	h.promMetrics.RequestCount.WithLabelValues(updateCarLabel).Inc()
	defer timer.ObserveDuration()

	carPrice := &pgtype.Numeric{}
	err := carPrice.UnmarshalJSON([]byte(strconv.FormatFloat(request.Car.Price, 'f', 2, 64)))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	req := postgres.UpdateCarParams{
		CarID:                 request.Car.CarId,
		BrandID:               request.Car.BrandId,
		Model:                 pgtype.Text{String: request.Car.Model, Valid: true},
		Year:                  pgtype.Int4{Int32: request.Car.Year, Valid: true},
		Price:                 *carPrice,
		Mileage:               pgtype.Int4{Int32: request.Car.Mileage, Valid: true},
		Color:                 pgtype.Text{String: request.Car.Color, Valid: true},
		Transmission:          pgtype.Text{String: request.Car.Transmission, Valid: true},
		FuelType:              pgtype.Text{String: request.Car.FuelType, Valid: true},
		EngineCapacity:        pgtype.Int4{Int32: request.Car.EngineCapacity, Valid: true},
		Description:           pgtype.Text{String: request.Car.Description, Valid: true},
		DealershipID:          pgtype.Text{String: request.Car.DealershipId, Valid: true},
		DealerID:              pgtype.Text{String: request.Car.DealerId, Valid: true},
		Status:                pgtype.Text{String: request.Car.Status, Valid: true},
		HorsePower:            pgtype.Int4{Int32: request.Car.HorsePower},
		Torque:                pgtype.Int4{Int32: request.Car.Torque},
		Title:                 pgtype.Text{String: request.Car.Title, Valid: true},
		SafetySpecifications:  pgtype.Array[string]{Elements: request.Car.SafetySpecifications},
		ComfortSpecifications: pgtype.Array[string]{Elements: request.Car.ComfortSpecifications},
		Ownership:             pgtype.Text{String: request.Car.Ownership, Valid: true},
	}

	updateCar, err := h.carService.UpdateCar(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(updateCarLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(updateCarLabel, "success").Inc()

	return &proto.UpdateCarResponse{
		CarId: updateCar,
	}, nil
}

var getOneCarLabel = "GetOneCar"

func (h *Handler) GetOneCar(ctx context.Context, request *proto.GetOneCarRequest) (*proto.GetOneCarResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getOneCarLabel))
	h.promMetrics.RequestCount.WithLabelValues(getOneCarLabel).Inc()
	defer timer.ObserveDuration()

	res, err := h.carService.GetCarById(ctx, request.CarId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getOneCarLabel, "error").Inc()
		return nil, status.Error(codes.NotFound, err.Error())
	}

	images, err := h.carService.ListCarImagesForCar(ctx, request.CarId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getOneCarLabel, "error").Inc()
	}

	brand, err := h.carService.GetCarBrandById(ctx, res.BrandID)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getOneCarLabel, "error").Inc()
	}

	extras, err := h.carService.ListExtraFeaturesForCar(ctx, request.CarId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getOneCarLabel, "error").Inc()
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getOneCarLabel, "success").Inc()

	carPrice, err := res.Price.Float64Value()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var carImages []*proto.CarImage
	var carExtras []*proto.CarExtraFeature
	var carBrand *proto.CarBrand

	for _, image := range images {
		carImages = append(carImages, &proto.CarImage{
			CarImageId: image.CarImageID,
			CarId:      image.CarID,
			Url:        image.ImageUrl.String,
			CreatedAt:  timestamppb.New(image.CreatedAt.Time),
			UpdatedAt:  timestamppb.New(image.UpdatedAt.Time),
		})
	}

	for _, extra := range extras {
		carExtras = append(carExtras, &proto.CarExtraFeature{
			CarExtraId: extra.CarExtraFeatureID,
			CarId:      extra.CarID,
			Name:       extra.Name.String,
			Value:      extra.Value.String,
			CreatedAt:  timestamppb.New(extra.CreatedAt.Time),
			UpdatedAt:  timestamppb.New(extra.UpdatedAt.Time),
		})
	}
	carBrand = &proto.CarBrand{
		BrandId:   brand.BrandID,
		Name:      brand.Name.String,
		LogoUrl:   brand.LogoUrl.String,
		CreatedAt: timestamppb.New(brand.CreatedAt.Time),
		UpdatedAt: timestamppb.New(brand.UpdatedAt.Time),
	}

	car := proto.Car{
		CarId:                     res.CarID,
		BrandId:                   res.BrandID,
		Model:                     res.Model.String,
		Year:                      res.Year.Int32,
		Price:                     carPrice.Float64,
		Mileage:                   res.Mileage.Int32,
		Color:                     res.Color.String,
		Transmission:              res.Transmission.String,
		FuelType:                  res.FuelType.String,
		EngineCapacity:            res.EngineCapacity.Int32,
		Description:               res.Description.String,
		DealershipId:              res.DealershipID.String,
		DealerId:                  res.DealerID.String,
		Status:                    res.Status.String,
		Title:                     res.Title.String,
		HorsePower:                res.HorsePower.Int32,
		Torque:                    res.Torque.Int32,
		TorqueRpm:                 res.TorqueRpm.Int32,
		SafetySpecifications:      res.SafetySpecifications.Elements,
		PerformanceSpecifications: res.PerformanceSpecifications.Elements,
		ComfortSpecifications:     res.ComfortSpecifications.Elements,
		Location:                  res.Location.String,
		Ownership:                 res.Ownership.String,
		CreatedAt:                 timestamppb.New(res.CreatedAt.Time),
		UpdatedAt:                 timestamppb.New(res.UpdatedAt.Time),
	}

	return &proto.GetOneCarResponse{
		Car: &proto.SingleCar{
			Car:    &car,
			Images: carImages,
			Extras: carExtras,
			Brand:  carBrand,
		},
	}, nil
}

var getCarsLabel = "GetCars"

func (h *Handler) GetCarsPaginated(ctx context.Context, request *proto.GetCarsPaginatedRequest) (*proto.GetCarsPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarsLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarsLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListCarsPagedParams{
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	res, err := h.carService.ListCarsPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarsLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	var cars []*proto.Car
	for _, car := range res {
		carPrice, err := car.Price.Float64Value()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		cars = append(cars, &proto.Car{
			CarId:          car.CarID,
			BrandId:        car.BrandID,
			Model:          car.Model.String,
			Year:           car.Year.Int32,
			Price:          carPrice.Float64,
			Mileage:        car.Mileage.Int32,
			Color:          car.Color.String,
			Transmission:   car.Transmission.String,
			FuelType:       car.FuelType.String,
			EngineCapacity: car.EngineCapacity.Int32,
			Description:    car.Description.String,
			DealershipId:   car.DealershipID.String,
			DealerId:       car.DealerID.String,
			CreatedAt:      timestamppb.New(car.CreatedAt.Time),
			UpdatedAt:      timestamppb.New(car.UpdatedAt.Time),
		})
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarsLabel, "success").Inc()

	return &proto.GetCarsPaginatedResponse{
		Cars:      cars,
		Size:      int32(len(cars)),
		TotalCars: -1,
	}, nil
}

var getCarsByDealershipIDLabel = "GetCarsByDealershipID"

func (h *Handler) GetCarsByDealershipIDPaginated(ctx context.Context, request *proto.GetCarsByDealershipIDPaginatedRequest) (*proto.GetCarsByDealershipIDPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarsByDealershipIDLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarsByDealershipIDLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListCarsByDealershipPagedParams{
		DealershipID: pgtype.Text{
			String: request.DealershipId,
			Valid:  true,
		},
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	res, err := h.carService.ListCarsByDealershipPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarsByDealershipIDLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	var cars []*proto.Car
	for _, car := range res {
		carPrice, err := car.Price.Float64Value()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		cars = append(cars, &proto.Car{
			CarId:          car.CarID,
			BrandId:        car.BrandID,
			Model:          car.Model.String,
			Year:           car.Year.Int32,
			Price:          carPrice.Float64,
			Mileage:        car.Mileage.Int32,
			Color:          car.Color.String,
			Transmission:   car.Transmission.String,
			FuelType:       car.FuelType.String,
			EngineCapacity: car.EngineCapacity.Int32,
			Description:    car.Description.String,
			DealershipId:   car.DealershipID.String,
			DealerId:       car.DealerID.String,
			CreatedAt:      timestamppb.New(car.CreatedAt.Time),
			UpdatedAt:      timestamppb.New(car.UpdatedAt.Time),
		})
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarsByDealershipIDLabel, "success").Inc()

	return &proto.GetCarsByDealershipIDPaginatedResponse{
		Cars:      cars,
		Size:      int32(len(cars)),
		TotalCars: -1,
	}, nil
}

var getCarsByBrandIDLabel = "GetCarsByBrandID"

func (h *Handler) GetCarsByBrandIDPaginated(ctx context.Context, request *proto.GetCarsByBrandIDPaginatedRequest) (*proto.GetCarsByCarBrandIDPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarsByBrandIDLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarsByBrandIDLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListCarsByBrandPagedParams{
		BrandID: request.CarBrandId,
		Offset:  request.Offset,
		Limit:   request.Limit,
	}

	res, err := h.carService.ListCarsByBrandPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarsByBrandIDLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	var cars []*proto.Car
	for _, car := range res {
		carPrice, err := car.Price.Float64Value()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		cars = append(cars, &proto.Car{
			CarId:          car.CarID,
			BrandId:        car.BrandID,
			Model:          car.Model.String,
			Year:           car.Year.Int32,
			Price:          carPrice.Float64,
			Mileage:        car.Mileage.Int32,
			Color:          car.Color.String,
			Transmission:   car.Transmission.String,
			FuelType:       car.FuelType.String,
			EngineCapacity: car.EngineCapacity.Int32,
			Description:    car.Description.String,
			DealershipId:   car.DealershipID.String,
			DealerId:       car.DealerID.String,
			CreatedAt:      timestamppb.New(car.CreatedAt.Time),
			UpdatedAt:      timestamppb.New(car.UpdatedAt.Time),
		})
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarsByBrandIDLabel, "success").Inc()

	return &proto.GetCarsByCarBrandIDPaginatedResponse{
		Cars:      cars,
		Size:      int32(len(cars)),
		TotalCars: -1,
	}, nil
}

var getCarsByDealerIDLabel = "GetCarsByDealerID"

func (h *Handler) GetCarsByDealerIDPaginated(ctx context.Context, request *proto.GetCarsByDealerIDPaginatedRequest) (*proto.GetCarsByDealerIDPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarsByDealerIDLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarsByDealerIDLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.ListCarsByDealerPagedParams{
		DealerID: pgtype.Text{String: request.DealerId, Valid: true},
		Offset:   request.Offset,
		Limit:    request.Limit,
	}

	res, err := h.carService.ListCarsByDealerPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarsByDealerIDLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	var cars []*proto.Car
	for _, car := range res {
		carPrice, err := car.Price.Float64Value()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		cars = append(cars, &proto.Car{
			CarId:                     car.CarID,
			BrandId:                   car.BrandID,
			Model:                     car.Model.String,
			Year:                      car.Year.Int32,
			Price:                     carPrice.Float64,
			Mileage:                   car.Mileage.Int32,
			Color:                     car.Color.String,
			Transmission:              car.Transmission.String,
			FuelType:                  car.FuelType.String,
			EngineCapacity:            car.EngineCapacity.Int32,
			Description:               car.Description.String,
			DealershipId:              car.DealershipID.String,
			DealerId:                  car.DealerID.String,
			Status:                    car.Status.String,
			Title:                     car.Title.String,
			IsFeatured:                car.IsFeatured.Bool,
			IsSold:                    car.IsSold.Bool,
			HorsePower:                car.HorsePower.Int32,
			Torque:                    car.Torque.Int32,
			TorqueRpm:                 car.TorqueRpm.Int32,
			SafetySpecifications:      car.SafetySpecifications.Elements,
			PerformanceSpecifications: car.PerformanceSpecifications.Elements,
			ComfortSpecifications:     car.ComfortSpecifications.Elements,
			Location:                  car.Location.String,
			Ownership:                 car.Ownership.String,
			CreatedAt:                 timestamppb.New(car.CreatedAt.Time),
			UpdatedAt:                 timestamppb.New(car.UpdatedAt.Time),
		})
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarsByDealerIDLabel, "success").Inc()

	return &proto.GetCarsByDealerIDPaginatedResponse{
		Cars:      cars,
		Size:      int32(len(cars)),
		TotalCars: -1,
	}, nil
}

var searchCarsLabel = "SearchCarsPaginated"

func (h *Handler) SearchCarsPaginated(ctx context.Context, request *proto.SearchCarsPaginatedRequest) (*proto.SearchCarsPaginatedResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(searchCarsLabel))
	h.promMetrics.RequestCount.WithLabelValues(searchCarsLabel).Inc()
	defer timer.ObserveDuration()

	req := postgres.SearchCarsPagedParams{
		Lower:  request.Query,
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	res, err := h.carService.SearchCarsPaged(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(searchCarsLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	var cars []*proto.Car
	for _, car := range res {
		carPrice, err := car.Price.Float64Value()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		cars = append(cars, &proto.Car{
			CarId:          car.CarID,
			BrandId:        car.BrandID,
			Model:          car.Model.String,
			Year:           car.Year.Int32,
			Price:          carPrice.Float64,
			Mileage:        car.Mileage.Int32,
			Color:          car.Color.String,
			Transmission:   car.Transmission.String,
			FuelType:       car.FuelType.String,
			EngineCapacity: car.EngineCapacity.Int32,
			Description:    car.Description.String,
			DealershipId:   car.DealershipID.String,
			DealerId:       car.DealerID.String,
			CreatedAt:      timestamppb.New(car.CreatedAt.Time),
			UpdatedAt:      timestamppb.New(car.UpdatedAt.Time),
		})
	}

	h.promMetrics.ResponseStatus.WithLabelValues(searchCarsLabel, "success").Inc()

	return &proto.SearchCarsPaginatedResponse{
		Cars:      cars,
		Size:      int32(len(cars)),
		TotalCars: -1,
	}, nil
}

var getCarByFieldLabel = "GetCarByField"

func (h *Handler) GetCarByField(ctx context.Context, request *proto.GetCarByFieldRequest) (*proto.GetCarByFieldResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarByFieldLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarByFieldLabel).Inc()
	defer timer.ObserveDuration()

	var filter map[string]dto.CarFilter

	for k, f := range request.Filter {
		filter[k] = dto.CarFilter{
			Value:     f.Value,
			Condition: f.Condition.String(),
		}
	}

	res, err := h.carService.GetCarByField(ctx, filter)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarByFieldLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	var cars []*proto.Car

	for _, car := range res {
		carPrice, err := car.Price.Float64Value()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		cars = append(cars, &proto.Car{
			CarId:          car.CarID,
			BrandId:        car.BrandID,
			Model:          car.Model.String,
			Year:           car.Year.Int32,
			Price:          carPrice.Float64,
			Mileage:        car.Mileage.Int32,
			Color:          car.Color.String,
			Title:          car.Title.String,
			Transmission:   car.Transmission.String,
			FuelType:       car.FuelType.String,
			EngineCapacity: car.EngineCapacity.Int32,
			Description:    car.Description.String,
			DealershipId:   car.DealershipID.String,
			DealerId:       car.DealerID.String,
			CreatedAt:      timestamppb.New(car.CreatedAt.Time),
			UpdatedAt:      timestamppb.New(car.UpdatedAt.Time),
		})
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarByFieldLabel, "success").Inc()

	return &proto.GetCarByFieldResponse{
		Cars: cars,
	}, nil
}

var getCarCountByFieldLabel = "GetCarCountByDealer"

func (h *Handler) GetCarByDealerCount(ctx context.Context, request *proto.GetCarByDealerCountRequest) (*proto.GetCarByDealerCountResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarCountByFieldLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarCountByFieldLabel).Inc()
	defer timer.ObserveDuration()

	req := pgtype.Text{
		String: request.DealerId,
		Valid:  true,
	}

	res, err := h.carService.CarByDealerCount(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarCountByFieldLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarCountByFieldLabel, "success").Inc()

	return &proto.GetCarByDealerCountResponse{
		TotalCars: res,
	}, nil
}

var getCarCountByDealershipLabel = "GetCarCountByDealership"

func (h *Handler) GetCarByDealershipCount(ctx context.Context, request *proto.GetCarByDealershipCountRequest) (*proto.GetCarByDealershipCountResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarCountByDealershipLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarCountByDealershipLabel).Inc()
	defer timer.ObserveDuration()

	req := pgtype.Text{
		String: request.DealershipId,
		Valid:  true,
	}

	res, err := h.carService.CarByDealershipCount(ctx, req)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarCountByDealershipLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarCountByDealershipLabel, "success").Inc()

	return &proto.GetCarByDealershipCountResponse{
		TotalCars: res,
	}, nil
}

var getCarCountByBrandLabel = "GetCarCountByBrand"

func (h *Handler) GetCarByBrandCount(ctx context.Context, request *proto.GetCarByBrandCountRequest) (*proto.GetCarByCarBrandCountResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(getCarCountByBrandLabel))
	h.promMetrics.RequestCount.WithLabelValues(getCarCountByBrandLabel).Inc()
	defer timer.ObserveDuration()

	res, err := h.carService.CarByBrandCount(ctx, request.CarBrandId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(getCarCountByBrandLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(getCarCountByBrandLabel, "success").Inc()

	return &proto.GetCarByCarBrandCountResponse{
		TotalCars: res,
	}, nil
}

var deleteCarLabel = "DeleteCar"

func (h *Handler) DeleteCar(ctx context.Context, request *proto.DeleteCarRequest) (*proto.DeleteCarResponse, error) {
	timer := prometheus.NewTimer(h.promMetrics.ResponseDuration.WithLabelValues(deleteCarLabel))
	h.promMetrics.RequestCount.WithLabelValues(deleteCarLabel).Inc()
	defer timer.ObserveDuration()

	id, err := h.carService.DeleteCar(ctx, request.CarId)
	if err != nil {
		h.promMetrics.ResponseStatus.WithLabelValues(deleteCarLabel, "error").Inc()
		return nil, status.Error(codes.Internal, err.Error())
	}

	h.promMetrics.ResponseStatus.WithLabelValues(deleteCarLabel, "success").Inc()

	return &proto.DeleteCarResponse{
		CarId: id,
	}, nil
}
