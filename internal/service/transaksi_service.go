package service

import (
	"context"
	"Rental/entity"
)

// usecase
type TransaksiUsecase interface {
	GetTransaksiByID(ctx context.Context, id int64) (*entity.Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]*entity.Transaksi, error)
	GetTransaksiByUserID(ctx context.Context, userID int64) ([]*entity.Transaksi, error) // Add this line
	GetProdukByID(ctx context.Context, id int64) (*entity.Produk, error) // Add this line

}

// repository
type TransaksiRepository interface {
	GetTransaksiByID(ctx context.Context, id int64) (*entity.Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]*entity.Transaksi, error)
	GetTransaksiByUserID(ctx context.Context, userID int64) ([]*entity.Transaksi, error) // Add this line
	GetProdukByID(ctx context.Context, id int64) (*entity.Produk, error) // Add this line
}

// service
type TransaksiService struct {
	repository TransaksiRepository
}

// newtransaksiservice
func NewTransaksiService(repository TransaksiRepository) *TransaksiService {
	return &TransaksiService{repository: repository}
}

// gettransaksibyid
func (s *TransaksiService) GetTransaksiByID(ctx context.Context, id int64) (*entity.Transaksi, error) {
	return s.repository.GetTransaksiByID(ctx, id)
}

// getalltransaksi
func (s *TransaksiService) GetAllTransaksi(ctx context.Context) ([]*entity.Transaksi, error) {
	return s.repository.GetAllTransaksi(ctx)
}

// GetTransaksiByUserID
func (s *TransaksiService) GetTransaksiByUserID(ctx context.Context, userID int64) ([]*entity.Transaksi, error) {
    return s.repository.GetTransaksiByUserID(ctx, userID)
}

// GetProdukByID
func (s *TransaksiService) GetProdukByID(ctx context.Context, id int64) (*entity.Produk, error) {
    return s.repository.GetProdukByID(ctx, id)
}