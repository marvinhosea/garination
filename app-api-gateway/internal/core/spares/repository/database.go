package repository

import (
	"context"
	"garination.com/gateway/internal/core/spares/ports"
	"garination.com/gateway/internal/platform/grpc/app-db/proto"
)

type sparePartRepositoryImpl struct {
	grpcClient proto.DatabaseServiceClient
}

func (s sparePartRepositoryImpl) GetSparePartByID(ctx context.Context, request *proto.GetSparePartByIDRequest) (*proto.GetSparePartByIDResponse, error) {
	res, err := s.grpcClient.GetSparePartByID(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) DeleteSparePart(ctx context.Context, request *proto.DeleteSparePartRequest) (*proto.DeleteSparePartResponse, error) {
	res, err := s.grpcClient.DeleteSparePart(ctx, request)
	if err != nil {
		return nil, err
	}

	return &proto.DeleteSparePartResponse{
		SparePartId: res.SparePartId,
	}, nil
}

func (s sparePartRepositoryImpl) FilterSparePartByBrandPaginated(ctx context.Context, request *proto.FilterSparePartByBrandPaginatedRequest) (*proto.FilterSparePartByBrandPaginatedResponse, error) {

	res, err := s.grpcClient.FilterSparePartByBrandPaginated(ctx, request)
	if err != nil {
		return nil, err
	}

	return &proto.FilterSparePartByBrandPaginatedResponse{
		SpareParts: res.SpareParts,
	}, nil
}

func (s sparePartRepositoryImpl) FilterSparePartByCategory(ctx context.Context, request *proto.FilterSparePartByCategoryPaginatedRequest) (*proto.FilterSparePartByCategoryPaginatedResponse, error) {
	res, err := s.grpcClient.FilterSparePartByCategory(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) FilterSparePartByCarModel(ctx context.Context, request *proto.FilterSparePartByCarModelPaginatedRequest) (*proto.FilterSparePartByCarModelPaginatedResponse, error) {
	res, err := s.grpcClient.FilterSparePartByCarModel(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) GetSparePartImageByID(ctx context.Context, request *proto.GetSparePartImageByIDRequest) (*proto.GetSparePartImageByIDResponse, error) {
	res, err := s.grpcClient.GetSparePartImageByID(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) InsertSparePart(ctx context.Context, request *proto.InsertSparePartRequest) (*proto.InsertSparePartResponse, error) {
	res, err := s.grpcClient.InsertSparePart(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) InsertSparePartImage(ctx context.Context, request *proto.InsertSparePartImageRequest) (*proto.InsertSparePartImageResponse, error) {
	res, err := s.grpcClient.InsertSparePartImage(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) ListSparePartImagesBySparePartPaged(ctx context.Context, request *proto.ListSparePartImagesBySparePartPagedRequest) (*proto.ListSparePartImagesBySparePartPagedResponse, error) {
	res, err := s.grpcClient.ListSparePartImagesBySparePartPaged(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) ListSparePartImagesPaged(ctx context.Context, request *proto.ListSparePartImagesPagedRequest) (*proto.ListSparePartImagesPagedResponse, error) {
	res, err := s.grpcClient.ListSparePartImagesPaged(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) ListSparePartsByDealerPaged(ctx context.Context, request *proto.ListSparePartsByDealerPagedRequest) (*proto.ListSparePartsByDealerPagedResponse, error) {
	res, err := s.grpcClient.ListSparePartsByDealerPaged(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) ListSparePartsByDealershipPaged(ctx context.Context, request *proto.ListSparePartsByDealershipPagedRequest) (*proto.ListSparePartsByDealershipPagedResponse, error) {
	res, err := s.grpcClient.ListSparePartsByDealershipPaged(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) ListSparePartsPaged(ctx context.Context, request *proto.ListSparePartsPagedRequest) (*proto.ListSparePartsPagedResponse, error) {
	res, err := s.grpcClient.ListSparePartsPaged(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) SearchSparePartsPaged(ctx context.Context, request *proto.SearchSparePartsPagedRequest) (*proto.SearchSparePartsPagedResponse, error) {
	res, err := s.grpcClient.SearchSparePartsPaged(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) UpdateSparePart(ctx context.Context, request *proto.UpdateSparePartRequest) (*proto.UpdateSparePartResponse, error) {
	res, err := s.grpcClient.UpdateSparePart(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s sparePartRepositoryImpl) UpdateSparePartImage(ctx context.Context, request *proto.UpdateSparePartImageRequest) (*proto.UpdateSparePartImageResponse, error) {
	res, err := s.grpcClient.UpdateSparePartImage(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewSparePartRepository(grpcClient proto.DatabaseServiceClient) ports.SpareRepository {
	return &sparePartRepositoryImpl{grpcClient: grpcClient}
}
