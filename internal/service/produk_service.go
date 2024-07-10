package service

import (
	"context"

	"Rental/entity"
)

// usecase
type ProdukUsecase interface {
	GetAllProduk(ctx context.Context) ([]entity.Produk, error)
	CreateProduk(ctx context.Context, produk entity.Produk) error
	UpdateProduk(ctx context.Context, produk entity.Produk) error
	DeleteProduk(ctx context.Context, id int) error
	GetProdukByLokasi(ctx context.Context, lokasiID int) ([]entity.Produk, error)
	GetProdukByPrice(ctx context.Context, price float64) ([]entity.Produk, error)
	SearchProdukByName(ctx context.Context, nameProduk string) ([]entity.Produk, error)
	GetMyProduk(ctx context.Context, OwnerID string) ([]entity.Produk, error)
	GetProdukByOwnerID(ctx context.Context, ownerID string) ([]entity.Produk, error)
}

// repository
type ProdukRepository interface {
	GetProduk(ctx context.Context) ([]entity.Produk, error)
	CreateProduk(ctx context.Context, produk entity.Produk) error
	UpdateProduk(ctx context.Context, produk entity.Produk) error
	DeleteProduk(ctx context.Context, id int) error
	GetProdukByLokasi(ctx context.Context, lokasiID int) ([]entity.Produk, error)
	GetProdukByPrice(ctx context.Context, price float64) ([]entity.Produk, error)
	SearchProdukByName(ctx context.Context, nameProduk string) ([]entity.Produk, error)
	GetMyProduk(ctx context.Context, OwnerID string) ([]entity.Produk, error)
	GetProdukByOwnerID(ctx context.Context, ownerID string) ([]entity.Produk, error)
}

// service
type ProdukService struct {
	repository ProdukRepository
}

// NewProdukService
func NewProdukService(repository ProdukRepository) ProdukUsecase {
	return &ProdukService{repository: repository}
}

// get all produk
func (s *ProdukService) GetAllProduk(ctx context.Context) ([]entity.Produk, error) {
	return s.repository.GetProduk(ctx)
}

// create produk
func (s *ProdukService) CreateProduk(ctx context.Context, produk entity.Produk) error {
	return s.repository.CreateProduk(ctx, produk)
}

// update produk
func (s *ProdukService) UpdateProduk(ctx context.Context, produk entity.Produk) error {
	return s.repository.UpdateProduk(ctx, produk)
}

// delete produk
func (s *ProdukService) DeleteProduk(ctx context.Context, id int) error {
	return s.repository.DeleteProduk(ctx, id)
}

// Get Produk by lokasi
func (s *ProdukService) GetProdukByLokasi(ctx context.Context, lokasiID int) ([]entity.Produk, error) {
	return s.repository.GetProdukByLokasi(ctx, lokasiID)
}

// Get produk by price
func (s *ProdukService) GetProdukByPrice(ctx context.Context, price float64) ([]entity.Produk, error) {
	return s.repository.GetProdukByPrice(ctx, price)
}

// search produk by name_produk
func (s *ProdukService) SearchProdukByName(ctx context.Context, nameProduk string) ([]entity.Produk, error) {
	return s.repository.SearchProdukByName(ctx, nameProduk)
}

// get my produk with jwt
func (s *ProdukService) GetMyProduk(ctx context.Context, OwnerID string) ([]entity.Produk, error) {
	return s.repository.GetMyProduk(ctx, OwnerID)
}

// get produk by owner id
func (s *ProdukService) GetProdukByOwnerID(ctx context.Context, ownerID string) ([]entity.Produk, error) {
	return s.repository.GetProdukByOwnerID(ctx, ownerID)
}

