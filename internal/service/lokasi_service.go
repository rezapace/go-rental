package service

import (
	"context"

	"Rental/entity"
)

// usecase
type LokasiUsecase interface {
	GetAllLokasi(ctx context.Context) (lokasi []entity.Lokasi, err error)
	CreateLokasi(ctx context.Context, lokasi entity.Lokasi) error
}


// repository
type LokasiRepository interface {
	GetAllLokasi(ctx context.Context) (lokasi []entity.Lokasi, err error)
	CreateLokasi(ctx context.Context, lokasi entity.Lokasi) error
}

// service
type LokasiService struct {
	repository LokasiRepository
}

// NewLokasiService
func NewLokasiService(repository LokasiRepository) LokasiUsecase {
	return &LokasiService{repository : repository}
}

// get all lokasi
func (s *LokasiService) GetAllLokasi(ctx context.Context) (lokasi []entity.Lokasi, err error) {
	return s.repository.GetAllLokasi(ctx)
}

// create lokasi
func (s *LokasiService) CreateLokasi(ctx context.Context, lokasi entity.Lokasi) error {
	return s.repository.CreateLokasi(ctx, lokasi)
}
