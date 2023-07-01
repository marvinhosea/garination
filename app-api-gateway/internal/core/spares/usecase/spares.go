package usecase

import (
	"context"
	"garination.com/gateway/internal/core/common"
	"garination.com/gateway/internal/core/spares/dto"
	"garination.com/gateway/internal/core/spares/model"
	"garination.com/gateway/internal/core/spares/ports"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
	"sync"
	"time"
)

type sparePartUsecaseImpl struct {
	sparePartRepository ports.SpareRepository
	wasabiRepository    common.WasabiRepo
}

const (
	defaultSpareDir = "spare-parts"
)

func protoTypeToSpareModel(p *proto.SparePart) model.SparePart {
	return model.SparePart{
		SparePartID:         p.SparePartId,
		Name:                p.Name,
		Description:         p.Description,
		Price:               p.Price,
		Used:                p.Used,
		CarModel:            p.CarModel,
		CarBrand:            p.CarBrand,
		OtherCompatibleCars: p.OtherCompatibleCars,
		CarYear:             p.CarYear,
		IsUniversal:         p.IsUniversal,
		Category:            p.Category,
		PartNumber:          p.PartNumber,
		DealershipID:        p.DealershipId,
		DealerID:            p.DealerId,
		CreatedAt:           p.CreatedAt.AsTime(),
		UpdatedAt:           p.UpdatedAt.AsTime(),
	}
}

func protoTypeToSpareModelList(p []*proto.SparePart) []model.SparePart {
	var spareParts []model.SparePart
	for _, v := range p {
		spareParts = append(spareParts, protoTypeToSpareModel(v))
	}
	return spareParts
}

func protoTypeToSpareImageModel(p *proto.SparePartImage) model.SparePartImage {
	return model.SparePartImage{
		SparePartImageID: p.SparePartImageId,
		SparePartID:      p.SparePartId,
		ImageURL:         p.ImageUrl,
		CreatedAt:        p.CreatedAt.AsTime(),
		UpdatedAt:        p.UpdatedAt.AsTime(),
	}
}

func protoTypeToSpareImageModelList(p []*proto.SparePartImage) []model.SparePartImage {
	var sparePartImages []model.SparePartImage
	for _, v := range p {
		sparePartImages = append(sparePartImages, protoTypeToSpareImageModel(v))
	}
	return sparePartImages
}

func (s sparePartUsecaseImpl) GetSparePartByID(ctx context.Context, request *dto.GetSparePartByIDRequest) (*dto.GetSparePartByIDResponse, error) {
	req := proto.GetSparePartByIDRequest{
		SparePartId: request.ID,
	}

	res, err := s.sparePartRepository.GetSparePartByID(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.GetSparePartByIDResponse{
		SparePart: protoTypeToSpareModel(res.SparePart),
	}, nil
}

func (s sparePartUsecaseImpl) DeleteSparePart(ctx context.Context, request *dto.DeleteSparePartRequest) (*dto.DeleteSparePartResponse, error) {
	req := proto.DeleteSparePartRequest{
		SparePartId: request.ID,
	}

	res, err := s.sparePartRepository.DeleteSparePart(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.DeleteSparePartResponse{
		SparePartId: res.SparePartId,
	}, nil
}

func (s sparePartUsecaseImpl) FilterSparePartByBrandPaginated(ctx context.Context, request *dto.FilterSparePartByBrandPaginatedRequest) (*dto.FilterSparePartByBrandPaginatedResponse, error) {
	req := proto.FilterSparePartByBrandPaginatedRequest{
		Query:  request.Query,
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	res, err := s.sparePartRepository.FilterSparePartByBrandPaginated(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.FilterSparePartByBrandPaginatedResponse{
		SpareParts:      protoTypeToSpareModelList(res.SpareParts),
		Size:            res.Size,
		TotalSpareParts: res.TotalSpareParts,
	}, nil
}

func (s sparePartUsecaseImpl) FilterSparePartByCategory(ctx context.Context, request *dto.FilterSparePartByCategoryPaginatedRequest) (*dto.FilterSparePartByCategoryPaginatedResponse, error) {
	req := proto.FilterSparePartByCategoryPaginatedRequest{
		Query:  request.Query,
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	res, err := s.sparePartRepository.FilterSparePartByCategory(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.FilterSparePartByCategoryPaginatedResponse{
		SpareParts: protoTypeToSpareModelList(res.SpareParts),
	}, nil
}

func (s sparePartUsecaseImpl) FilterSparePartByCarModel(ctx context.Context, request *dto.FilterSparePartByCarModelPaginatedRequest) (*dto.FilterSparePartByCarModelPaginatedResponse, error) {
	req := proto.FilterSparePartByCarModelPaginatedRequest{
		Query:  request.Query,
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	res, err := s.sparePartRepository.FilterSparePartByCarModel(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.FilterSparePartByCarModelPaginatedResponse{
		SpareParts: protoTypeToSpareModelList(res.SpareParts),
	}, nil
}

func (s sparePartUsecaseImpl) GetSparePartImageByID(ctx context.Context, request *dto.GetSparePartImageByIDRequest) (*dto.GetSparePartImageByIDResponse, error) {
	req := proto.GetSparePartImageByIDRequest{
		SparePartImageId: request.ID,
	}

	res, err := s.sparePartRepository.GetSparePartImageByID(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.GetSparePartImageByIDResponse{
		SparePartImage: protoTypeToSpareImageModel(res.SparePartImage),
	}, nil
}

func (s sparePartUsecaseImpl) InsertSparePart(ctx context.Context, request *dto.InsertSparePartRequest) (*dto.InsertSparePartResponse, error) {
	req := proto.InsertSparePartRequest{
		Name:                request.Name,
		Description:         request.Description,
		Price:               request.Price,
		Used:                request.Used,
		CarModel:            "",
		CarBrand:            "",
		OtherCompatibleCars: nil,
		CarYear:             0,
		IsUniversal:         false,
		Category:            request.Category,
		PartNumber:          request.PartNumber,
		DealershipId:        request.DealershipID,
		DealerId:            request.DealerID,
	}

	res, err := s.sparePartRepository.InsertSparePart(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.InsertSparePartResponse{
		SparePartId: res.SparePartId,
	}, nil
}

func (s sparePartUsecaseImpl) InsertSparePartImage(ctx context.Context, request *dto.InsertSparePartImageRequest) (*dto.InsertSparePartImageResponse, error) {

	var wg sync.WaitGroup
	wg.Add(len(request.Images))
	var errChan = make(chan error, 1)
	var linksChan = make(chan model.SparePartImage, len(request.Images))
	var images = make([]model.SparePartImage, 0, len(request.Images))

	for _, image := range request.Images {
		imgVal := image
		go func() {
			defer wg.Done()
			url, err := s.wasabiRepository.UploadImage(ctx, defaultSpareDir, imgVal)
			if err != nil {
				errChan <- err
				return
			}

			req := proto.InsertSparePartImageRequest{
				SparePartId: request.SparePartID,
				ImageUrl:    url,
			}

			res, err := s.sparePartRepository.InsertSparePartImage(ctx, &req)
			if err != nil {
				errChan <- err
				return
			}

			linksChan <- model.SparePartImage{
				SparePartImageID: res.SparePartImageId,
				SparePartID:      request.SparePartID,
				ImageURL:         url,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}

		}()
	}

	go func() {
		wg.Wait()
		close(linksChan)
		close(errChan)
	}()

	for {
		select {
		case err := <-errChan:
			return nil, err
		case link := <-linksChan:
			images = append(images, link)
		}
		if len(images) == len(request.Images) {
			break
		}
	}

	return &dto.InsertSparePartImageResponse{
		Images: images,
	}, nil
}

func (s sparePartUsecaseImpl) ListSparePartImagesBySparePartPaged(ctx context.Context, request *dto.ListSparePartImagesBySparePartPagedRequest) (*dto.ListSparePartImagesBySparePartPagedResponse, error) {
	req := proto.ListSparePartImagesBySparePartPagedRequest{
		SparePartId: request.SparePartId,
		Offset:      request.Offset,
		Limit:       request.Limit,
	}

	res, err := s.sparePartRepository.ListSparePartImagesBySparePartPaged(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.ListSparePartImagesBySparePartPagedResponse{
		SparePartImages: protoTypeToSpareImageModelList(res.SparePartImages),
	}, nil
}

func (s sparePartUsecaseImpl) ListSparePartImagesPaged(ctx context.Context, request *dto.ListSparePartImagesPagedRequest) (*dto.ListSparePartImagesPagedResponse, error) {
	req := proto.ListSparePartImagesPagedRequest{
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	res, err := s.sparePartRepository.ListSparePartImagesPaged(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.ListSparePartImagesPagedResponse{
		SparePartImages: protoTypeToSpareImageModelList(res.SparePartImages),
	}, nil
}

func (s sparePartUsecaseImpl) ListSparePartsByDealerPaged(ctx context.Context, request *dto.ListSparePartsByDealerPagedRequest) (*dto.ListSparePartsByDealerPagedResponse, error) {
	req := proto.ListSparePartsByDealerPagedRequest{
		DealerId: request.DealerId,
		Offset:   request.Offset,
		Limit:    request.Limit,
	}

	res, err := s.sparePartRepository.ListSparePartsByDealerPaged(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.ListSparePartsByDealerPagedResponse{
		SpareParts: protoTypeToSpareModelList(res.SpareParts),
	}, nil
}

func (s sparePartUsecaseImpl) ListSparePartsByDealershipPaged(ctx context.Context, request *dto.ListSparePartsByDealershipPagedRequest) (*dto.ListSparePartsByDealershipPagedResponse, error) {
	req := proto.ListSparePartsByDealershipPagedRequest{
		DealershipId: request.DealershipId,
		Offset:       request.Offset,
		Limit:        request.Limit,
	}

	res, err := s.sparePartRepository.ListSparePartsByDealershipPaged(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.ListSparePartsByDealershipPagedResponse{
		SpareParts: protoTypeToSpareModelList(res.SpareParts),
	}, nil
}

func (s sparePartUsecaseImpl) ListSparePartsPaged(ctx context.Context, request *dto.ListSparePartsPagedRequest) (*dto.ListSparePartsPagedResponse, error) {
	req := proto.ListSparePartsPagedRequest{
		Offset: request.Offset,
		Limit:  request.Limit,
	}

	res, err := s.sparePartRepository.ListSparePartsPaged(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.ListSparePartsPagedResponse{
		SpareParts: protoTypeToSpareModelList(res.SpareParts),
	}, nil
}

func (s sparePartUsecaseImpl) SearchSparePartsPaged(ctx context.Context, request *dto.SearchSparePartsPagedRequest) (*dto.SearchSparePartsPagedResponse, error) {
	req := proto.SearchSparePartsPagedRequest{
		Offset: request.Offset,
		Limit:  request.Limit,
		Query:  request.Query,
	}

	res, err := s.sparePartRepository.SearchSparePartsPaged(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.SearchSparePartsPagedResponse{
		SpareParts: protoTypeToSpareModelList(res.SpareParts),
	}, nil
}

func (s sparePartUsecaseImpl) UpdateSparePart(ctx context.Context, request *dto.UpdateSparePartRequest) (*dto.UpdateSparePartResponse, error) {
	req := proto.UpdateSparePartRequest{
		SparePartId:         request.SparePartId,
		Name:                request.Name,
		Description:         request.Description,
		Price:               request.Price,
		Used:                request.Used,
		CarModel:            request.CarModel,
		CarBrand:            request.CarBrand,
		OtherCompatibleCars: request.OtherCompatibleCars,
		CarYear:             request.CarYear,
		IsUniversal:         request.IsUniversal,
		Category:            request.Category,
		PartNumber:          request.PartNumber,
	}

	res, err := s.sparePartRepository.UpdateSparePart(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateSparePartResponse{
		SparePartId: res.SparePartId,
	}, nil
}

func (s sparePartUsecaseImpl) UpdateSparePartImage(ctx context.Context, request *dto.UpdateSparePartImageRequest) (*dto.UpdateSparePartImageResponse, error) {
	req := proto.UpdateSparePartImageRequest{
		SparePartId: request.SparePartId,
	}

	url, err := s.wasabiRepository.UploadImage(ctx, defaultSpareDir, request.Image)
	if err != nil {
		return nil, err
	}

	req.ImageUrl = url

	res, err := s.sparePartRepository.UpdateSparePartImage(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateSparePartImageResponse{
		SparePartImageId: res.SparePartImageId,
	}, nil
}

func NewSparePartUsecase(sparePartRepository ports.SpareRepository) ports.SpareUseCase {
	return &sparePartUsecaseImpl{
		sparePartRepository: sparePartRepository,
	}
}
