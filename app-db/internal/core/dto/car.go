package dto

import (
	"time"
)

type CarImage struct {
	CarImageID string
	CarID      string
	ImageUrl   string
	CreatedAt  time.Time
}

type CarImagesParam struct {
	CarImages []CarImage
}

type GetCarsPaginatedRequest struct {
	Offset int32
	Limit  int32
	SortBy int32
}

type CarFilter struct {
	Value     string
	Condition string
	Offset    int32
	Limit     int32
}
