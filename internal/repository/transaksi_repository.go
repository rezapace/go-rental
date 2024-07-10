package repository

import (
	"context"

	"Rental/entity"

	"gorm.io/gorm"
)

type TransaksiRepository struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) *TransaksiRepository {
	return &TransaksiRepository{db: db}
}

// gettransaksibyid
func (r *TransaksiRepository) GetTransaksiByID(ctx context.Context, id int64) (*entity.Transaksi, error) {
	var transaksi entity.Transaksi
	if err := r.db.WithContext(ctx).First(&transaksi, id).Error; err != nil {
		return nil, err
	}
	return &transaksi, nil
}

// getalltransaksi
func (r *TransaksiRepository) GetAllTransaksi(ctx context.Context) ([]*entity.Transaksi, error) {
	var transaksi []*entity.Transaksi
	if err := r.db.WithContext(ctx).Table("transaksi").Find(&transaksi).Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}

// gettransaksibyuserid
func (r *TransaksiRepository) GetTransaksiByUserID(ctx context.Context, userID int64) ([]*entity.Transaksi, error) {
    var transaksi []*entity.Transaksi
    if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&transaksi).Error; err != nil {
        return nil, err
    }
    return transaksi, nil
}

// getprodukbyid
func (r *TransaksiRepository) GetProdukByID(ctx context.Context, id int64) (*entity.Produk, error) {
    var produk entity.Produk
    if err := r.db.WithContext(ctx).First(&produk, id).Error; err != nil {
        return nil, err
    }
    return &produk, nil
}