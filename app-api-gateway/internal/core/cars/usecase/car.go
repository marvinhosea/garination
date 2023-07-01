package usecase

import (
	"context"
	"garination.com/gateway/internal/core/cars/dto"
	"garination.com/gateway/internal/core/cars/model"
	"garination.com/gateway/internal/core/cars/ports"
	"garination.com/gateway/internal/core/common"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mime/multipart"
	"sync"
)

type usecaseImpl struct {
	dbRepo     ports.CarDatabaseRepo
	wasabiRepo common.WasabiRepo
}

const (
	s3folder = "cars"
)

func (u usecaseImpl) InsertCarBrand(ctx context.Context, in *dto.CreateCarBrandRequest) (*dto.CreateCarBrandResponse, error) {
	carBrandModel := &proto.InsertCarBrandRequest{
		CarBrand: &proto.CarBrand{
			Name: in.Name,
		},
	}

	// upload image
	imageUrl, err := u.wasabiRepo.UploadImage(ctx, s3folder, in.Logo)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to upload image: "+err.Error())
	}

	carBrandModel.CarBrand.LogoUrl = imageUrl

	carBrandResponse, err := u.dbRepo.InsertCarBrand(ctx, carBrandModel)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCarBrandResponse{
		BrandID: carBrandResponse.CarBrandId,
		Logo:    imageUrl,
	}, nil
}

func (u usecaseImpl) UpdateCarBrand(ctx context.Context, in *dto.UpdateCarBrandRequest) (*dto.UpdateCarBrandResponse, error) {
	brandUpdateModel := &proto.UpdateCarBrandRequest{
		CarBrand: &proto.CarBrand{
			BrandId: in.BrandID,
			Name:    in.Name,
			LogoUrl: in.LogoURL,
		},
	}

	brandUpdateResponse, err := u.dbRepo.UpdateCarBrand(ctx, brandUpdateModel)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateCarBrandResponse{
		BrandID: brandUpdateResponse.CarBrandId,
	}, nil
}

func (u usecaseImpl) GetCarBrandByID(ctx context.Context, in *dto.GetCarBrandByIDRequest) (*dto.GetCarBrandByIDResponse, error) {
	carBrandsResponse, err := u.dbRepo.GetCarBrandByID(ctx, &proto.GetCarBrandByIDRequest{CarBrandId: in.BrandID})
	if err != nil {
		return nil, err
	}

	return &dto.GetCarBrandByIDResponse{
		BrandID:   carBrandsResponse.CarBrand.BrandId,
		Name:      carBrandsResponse.CarBrand.Name,
		LogoURL:   carBrandsResponse.CarBrand.LogoUrl,
		CreatedAt: carBrandsResponse.CarBrand.CreatedAt.AsTime(),
		UpdatedAt: carBrandsResponse.CarBrand.UpdatedAt.AsTime(),
	}, nil
}

func (u usecaseImpl) GetCarBrandsPaginated(ctx context.Context, in *dto.GetCarBrandsPaginatedRequest) (*dto.GetCarBrandsPaginatedResponse, error) {
	carBrandsResponse, err := u.dbRepo.GetCarBrandsPaginated(ctx, &proto.GetCarBrandsPaginatedRequest{
		Offset: in.Offset,
		Limit:  in.Limit,
	})
	if err != nil {
		return nil, err
	}

	var carBrands []dto.GetCarBrandByIDResponse
	for _, carBrand := range carBrandsResponse.CarBrands {
		carBrands = append(carBrands, dto.GetCarBrandByIDResponse{
			BrandID:   carBrand.BrandId,
			Name:      carBrand.Name,
			LogoURL:   carBrand.LogoUrl,
			CreatedAt: carBrand.CreatedAt.AsTime(),
			UpdatedAt: carBrand.UpdatedAt.AsTime(),
		})
	}

	return &dto.GetCarBrandsPaginatedResponse{
		Brands: carBrands,
		Total:  carBrandsResponse.TotalCars,
	}, nil
}

func (u usecaseImpl) DeleteCarBrand(ctx context.Context, in *dto.DeleteCarBrandRequest) (*dto.DeleteCarBrandResponse, error) {
	_, err := u.dbRepo.DeleteCarBrand(ctx, &proto.DeleteCarBrandRequest{CarBrandId: in.BrandID})
	if err != nil {
		return nil, err
	}

	return &dto.DeleteCarBrandResponse{}, nil
}

func (u usecaseImpl) InsertCarExtraFeature(ctx context.Context, in *dto.CreateCarExtraFeatureRequest) (*dto.CreateCarExtraFeatureResponse, error) {
	carExtraFeatureModel := &proto.InsertCarExtraFeatureRequest{
		CarExtraFeature: &proto.CarExtraFeature{
			CarId: in.CarID,
			Name:  in.Name,
			Value: in.Value,
		},
	}

	carExtraFeatureResponse, err := u.dbRepo.InsertCarExtraFeature(ctx, carExtraFeatureModel)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCarExtraFeatureResponse{
		ExtraFeatureID: carExtraFeatureResponse.CarExtraFeatureId,
	}, nil
}

func (u usecaseImpl) UpdateCarExtraFeature(ctx context.Context, in *dto.UpdateCarExtraFeatureRequest) (*dto.UpdateCarExtraFeatureResponse, error) {
	carExtraFeatureUpdateModel := &proto.UpdateCarExtraFeatureRequest{
		CarExtraFeature: &proto.CarExtraFeature{
			CarExtraId: in.ExtraFeatureID,
			Name:       in.Name,
			Value:      in.Value,
		},
	}

	carExtraFeatureUpdateResponse, err := u.dbRepo.UpdateCarExtraFeature(ctx, carExtraFeatureUpdateModel)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateCarExtraFeatureResponse{
		ExtraFeatureID: carExtraFeatureUpdateResponse.CarExtraFeatureId,
	}, nil
}

func (u usecaseImpl) GetCarExtraFeaturePaginated(ctx context.Context, in *dto.GetCarExtraFeaturesPaginatedRequest) (*dto.GetCarExtraFeaturesPaginatedResponse, error) {
	carExtraFeaturesResponse, err := u.dbRepo.GetCarExtraFeaturePaginated(ctx, &proto.GetCarExtraFeaturesPaginatedRequest{
		Offset: in.Offset,
		Limit:  in.Limit,
	})
	if err != nil {
		return nil, err
	}

	var carExtraFeatures []model.CarExtraFeature
	for _, carExtraFeature := range carExtraFeaturesResponse.CarExtraFeatures {
		carExtraFeatures = append(carExtraFeatures, model.CarExtraFromProto(carExtraFeature))
	}

	return &dto.GetCarExtraFeaturesPaginatedResponse{
		ExtraFeatures: carExtraFeatures,
		Total:         carExtraFeaturesResponse.TotalCars,
	}, nil
}

func (u usecaseImpl) DeleteCarExtraFeature(ctx context.Context, in *dto.DeleteCarExtraFeatureRequest) (*dto.DeleteCarExtraFeatureResponse, error) {
	_, err := u.dbRepo.DeleteCarExtraFeature(ctx, &proto.DeleteCarExtraFeatureRequest{CarExtraFeatureId: in.ExtraFeatureID})
	if err != nil {
		return nil, err
	}

	return &dto.DeleteCarExtraFeatureResponse{}, nil
}

func (u usecaseImpl) CreateCarImage(ctx context.Context, in *dto.CreateCarImageRequest) (*dto.CreateCarImageResponse, error) {
	carImageModel := &proto.CreateCarImageRequest{}

	// Create a channel to receive upload results
	results := make(chan string)

	// Create a WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// Set the Goroutine count to the number of images
	wg.Add(len(in.Images))

	// Process each car image concurrently
	for _, carImage := range in.Images {
		go func(image *multipart.FileHeader) {
			defer wg.Done()

			// Upload the image asynchronously
			res, err := u.wasabiRepo.UploadImage(ctx, "car-images", image)
			if err != nil {
				results <- err.Error()
				return
			}

			// Send the result through the results channel
			results <- res
		}(carImage)
	}

	// Start a Goroutine to collect the results
	go func() {
		// Wait for all Goroutines to finish
		wg.Wait()

		// Close the results channel
		close(results)
	}()

	// Collect the results from the channel
	var urls []string
	for result := range results {
		urls = append(urls, result)
	}

	// Build the carImageModel
	for _, url := range urls {
		carImageModel.CarImage = append(carImageModel.CarImage, &proto.CarImage{
			CarId: in.CarID,
			Url:   url,
		})
	}

	// Call the database repository to create the car image
	carImageResponse, err := u.dbRepo.CreateCarImage(ctx, carImageModel)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCarImageResponse{
		ImageID: carImageResponse.CarImageIds,
	}, nil
}

func (u usecaseImpl) UpdateCarImage(ctx context.Context, in *dto.UpdateCarImageRequest) (*dto.UpdateCarImageResponse, error) {
	carImageModel := &proto.UpdateCarImageRequest{
		CarImage: &proto.CarImage{
			CarImageId: in.ImageID,
			Url:        in.URL,
		},
	}

	carImageResponse, err := u.dbRepo.UpdateCarImage(ctx, carImageModel)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateCarImageResponse{
		ImageID: carImageResponse.CarImageId,
	}, nil
}

func (u usecaseImpl) GetCarImagePaginated(ctx context.Context, in *dto.GetCarImagesPaginatedRequest) (*dto.GetCarImagesPaginatedResponse, error) {
	carImagesResponse, err := u.dbRepo.GetCarImagePaginated(ctx, &proto.GetCarImagesPaginatedRequest{
		Offset: in.Offset,
		Limit:  in.Limit,
	})
	if err != nil {
		return nil, err
	}

	var carImages []model.CarImage
	for _, carImage := range carImagesResponse.CarImages {
		carImages = append(carImages, model.CarImageFromProto(carImage))
	}

	return &dto.GetCarImagesPaginatedResponse{
		Images:    carImages,
		TotalCars: carImagesResponse.TotalCars,
	}, nil
}

func (u usecaseImpl) DeleteCarImage(ctx context.Context, in *dto.DeleteCarImageRequest) (*dto.DeleteCarImageResponse, error) {
	_, err := u.dbRepo.DeleteCarImage(ctx, &proto.DeleteCarImageRequest{CarImageId: in.ImageID})
	if err != nil {
		return nil, err
	}

	return &dto.DeleteCarImageResponse{}, nil
}

func (u usecaseImpl) InsertCar(ctx context.Context, in *dto.InsertCarRequest) (*dto.InsertCarResponse, error) {
	carModel := &proto.InsertCarRequest{
		Car: &proto.Car{
			BrandId:                   in.BrandID,
			Model:                     in.Model,
			Year:                      in.Year,
			Price:                     in.Price,
			Mileage:                   in.Mileage,
			Color:                     in.Color,
			Transmission:              in.Transmission,
			FuelType:                  in.FuelType,
			EngineCapacity:            in.EngineCapacity,
			Description:               in.Description,
			DealershipId:              in.DealershipID,
			DealerId:                  in.DealerID,
			Title:                     in.Title,
			HorsePower:                in.HorsePower,
			Torque:                    in.Torque,
			TorqueRpm:                 in.TorqueRpm,
			SafetySpecifications:      in.SafetySpecifications,
			PerformanceSpecifications: in.PerformanceSpecifications,
			ComfortSpecifications:     in.ComfortSpecifications,
			Location:                  in.Location,
			Ownership:                 in.Ownership,
		},
	}

	carResponse, err := u.dbRepo.InsertCar(ctx, carModel)
	if err != nil {
		return nil, err
	}

	return &dto.InsertCarResponse{
		CarID: carResponse.CarId,
	}, nil
}

func (u usecaseImpl) UpdateCar(ctx context.Context, in *dto.UpdateCarRequest) (*dto.UpdateCarResponse, error) {
	carModel := &proto.UpdateCarRequest{
		Car: &proto.Car{
			CarId:          in.CarID,
			BrandId:        in.BrandID,
			Model:          in.Model,
			Year:           in.Year,
			Price:          in.Price,
			Mileage:        in.Mileage,
			Color:          in.Color,
			Transmission:   in.Transmission,
			FuelType:       in.FuelType,
			EngineCapacity: in.EngineCapacity,
			Description:    in.Description,
		},
	}

	carResponse, err := u.dbRepo.UpdateCar(ctx, carModel)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateCarResponse{
		CarID: carResponse.CarId,
	}, nil
}

func (u usecaseImpl) GetOneCar(ctx context.Context, in *dto.GetOneCarRequest) (*dto.GetOneCarResponse, error) {
	carResponse, err := u.dbRepo.GetOneCar(ctx, &proto.GetOneCarRequest{CarId: in.CarID})
	if err != nil {
		return nil, err
	}

	var carImages []model.CarImage
	var carFeatures []model.CarExtraFeature

	for _, carImage := range carResponse.Car.Images {
		carImages = append(carImages, model.CarImageFromProto(carImage))
	}

	for _, carFeature := range carResponse.Car.Extras {
		carFeatures = append(carFeatures, model.CarExtraFromProto(carFeature))
	}

	return &dto.GetOneCarResponse{
		Car:           model.CarFromProto(carResponse.Car.Car),
		Images:        carImages,
		ExtraFeatures: carFeatures,
	}, nil
}

func (u usecaseImpl) GetCarsPaginated(ctx context.Context, in *dto.GetCarsPaginatedRequest) (*dto.GetCarsPaginatedResponse, error) {
	carsResponse, err := u.dbRepo.GetCarsPaginated(ctx, &proto.GetCarsPaginatedRequest{
		Offset: in.Offset,
		Limit:  in.Limit,
	})
	if err != nil {
		return nil, err
	}

	var cars []model.Car
	for _, car := range carsResponse.Cars {
		cars = append(cars, model.CarFromProto(car))
	}

	return &dto.GetCarsPaginatedResponse{
		Cars: cars,
	}, nil
}

func (u usecaseImpl) GetCarsByDealershipIDPaginated(ctx context.Context, in *dto.GetCarsByDealershipIDPaginatedRequest) (*dto.GetCarsByDealershipIDPaginatedResponse, error) {
	getDealershipReq := &proto.GetCarsByDealershipIDPaginatedRequest{
		DealershipId: in.DealershipID,
	}

	getDealershipRes, err := u.dbRepo.GetCarsByDealershipIDPaginated(ctx, getDealershipReq)
	if err != nil {
		return nil, err
	}

	var cars []model.Car
	for _, car := range getDealershipRes.Cars {
		cars = append(cars, model.CarFromProto(car))
	}

	return &dto.GetCarsByDealershipIDPaginatedResponse{
		Cars: cars,
	}, nil
}

func (u usecaseImpl) GetCarsByBrandIDPaginated(ctx context.Context, in *dto.GetCarsByBrandIDPaginatedRequest) (*dto.GetCarsByCarBrandIDPaginatedResponse, error) {
	getBrandReq := &proto.GetCarsByBrandIDPaginatedRequest{
		CarBrandId: in.BrandID,
	}

	getBrandRes, err := u.dbRepo.GetCarsByBrandIDPaginated(ctx, getBrandReq)
	if err != nil {
		return nil, err
	}

	var cars []model.Car
	for _, car := range getBrandRes.Cars {
		cars = append(cars, model.CarFromProto(car))
	}

	return &dto.GetCarsByCarBrandIDPaginatedResponse{
		Cars:  cars,
		Total: getBrandRes.TotalCars,
	}, nil
}

func (u usecaseImpl) GetCarsByDealerIDPaginated(ctx context.Context, in *dto.GetCarsByDealerIDPaginatedRequest) (*dto.GetCarsByDealerIDPaginatedResponse, error) {
	getDealerReq := &proto.GetCarsByDealerIDPaginatedRequest{
		DealerId: in.DealerID,
	}

	getDealerRes, err := u.dbRepo.GetCarsByDealerIDPaginated(ctx, getDealerReq)
	if err != nil {
		return nil, err
	}

	var cars []model.Car
	for _, car := range getDealerRes.Cars {
		cars = append(cars, model.CarFromProto(car))
	}

	return &dto.GetCarsByDealerIDPaginatedResponse{
		Cars:  cars,
		Total: getDealerRes.TotalCars,
	}, nil
}

func (u usecaseImpl) SearchCarsPaginated(ctx context.Context, in *dto.SearchCarsPaginatedRequest) (*dto.SearchCarsPaginatedResponse, error) {
	searchReq := &proto.SearchCarsPaginatedRequest{
		Query:  in.Search,
		Offset: in.Offset,
		Limit:  in.Limit,
	}

	searchRes, err := u.dbRepo.SearchCarsPaginated(ctx, searchReq)
	if err != nil {
		return nil, err
	}

	var cars []model.Car
	for _, car := range searchRes.Cars {
		cars = append(cars, model.CarFromProto(car))
	}

	return &dto.SearchCarsPaginatedResponse{
		Cars:  cars,
		Total: searchRes.TotalCars,
	}, nil
}

func (u usecaseImpl) GetCarsByField(ctx context.Context, in *dto.GetCarsByFieldRequest) (*dto.GetCarsByFieldResponse, error) {
	var filters map[string]*proto.CarFilter

	for _, filter := range in.Filters {
		var condition proto.Condition
		switch filter.Condition {
		case "equal":
			condition = proto.Condition_EQUAL
		case "not_equal":
			condition = proto.Condition_NOT_EQUAL
		case "greater_than":
			condition = proto.Condition_GREATER_THAN
		case "less_than":
			condition = proto.Condition_LESS_THAN
		case "greater_than_or_equal":
			condition = proto.Condition_GREATER_THAN_OR_EQUAL
		case "less_than_or_equal":
			condition = proto.Condition_LESS_THAN_OR_EQUAL
		case "like":
			condition = proto.Condition_LIKE
		case "not_like":
			condition = proto.Condition_NOT_LIKE
		case "in":
			condition = proto.Condition_IN
		default:
			return nil, status.Error(codes.InvalidArgument, "invalid condition")
		}

		filters[filter.Field] = &proto.CarFilter{
			Condition: condition,
			Value:     filter.Value,
		}
	}

	req := &proto.GetCarByFieldRequest{
		Filter: filters,
		Offset: in.Offset,
		Limit:  in.Limit,
	}

	res, err := u.dbRepo.GetCarByField(ctx, req)
	if err != nil {
		return nil, err
	}

	var cars []model.Car

	for _, car := range res.Cars {
		cars = append(cars, model.CarFromProto(car))
	}

	return &dto.GetCarsByFieldResponse{
		Cars:  cars,
		Total: res.TotalCars,
	}, nil
}

func (u usecaseImpl) GetCarByDealerCount(ctx context.Context, in *dto.GetCarByDealerCountRequest) (*dto.GetCarByDealerCountResponse, error) {
	res, err := u.dbRepo.GetCarByDealerCount(ctx, &proto.GetCarByDealerCountRequest{DealerId: in.DealerID})
	if err != nil {
		return nil, err
	}

	return &dto.GetCarByDealerCountResponse{
		Total: res.TotalCars,
	}, nil
}

func (u usecaseImpl) GetCarByDealershipCount(ctx context.Context, in *dto.GetCarByDealershipCountRequest) (*dto.GetCarByDealershipCountResponse, error) {
	res, err := u.dbRepo.GetCarByDealershipCount(ctx, &proto.GetCarByDealershipCountRequest{DealershipId: in.DealershipID})
	if err != nil {
		return nil, err
	}

	return &dto.GetCarByDealershipCountResponse{
		Total: res.TotalCars,
	}, nil
}

func (u usecaseImpl) GetCarByBrandCount(ctx context.Context, in *dto.GetCarByBrandCountRequest) (*dto.GetCarByCarBrandCountResponse, error) {
	res, err := u.dbRepo.GetCarByBrandCount(ctx, &proto.GetCarByBrandCountRequest{CarBrandId: in.BrandID})
	if err != nil {
		return nil, err
	}

	return &dto.GetCarByCarBrandCountResponse{
		Total: res.TotalCars,
	}, nil
}

func (u usecaseImpl) DeleteCar(ctx context.Context, in *dto.DeleteCarRequest) (*dto.DeleteCarResponse, error) {
	res, err := u.dbRepo.DeleteCar(ctx, &proto.DeleteCarRequest{CarId: in.CarID})
	if err != nil {
		return nil, err
	}

	return &dto.DeleteCarResponse{
		CarID: res.CarId,
	}, nil
}

func NewUsecase(dbRepo ports.CarDatabaseRepo, wasabiRepo common.WasabiRepo) ports.CarUsercase {
	return &usecaseImpl{
		dbRepo:     dbRepo,
		wasabiRepo: wasabiRepo,
	}
}
