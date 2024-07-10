package repository

import (
	"context"

	"Rental/entity"

	"gorm.io/gorm"
)

// repository
type LokasiRepository struct {
	db *gorm.DB
}

// NewLokasiRepository
func NewLokasiRepository(db *gorm.DB) *LokasiRepository {
	return &LokasiRepository{db: db}
}

// get all Lokasi
func (r *LokasiRepository) GetAllLokasi(ctx context.Context) ([]entity.Lokasi, error) {
	var lokasi []entity.Lokasi
	err := r.db.Table("lokasi").Find(&lokasi).Error
	if err != nil {
		return nil, err
	}
	return lokasi, nil
}

// create lokasi
func (r *LokasiRepository) CreateLokasi(ctx context.Context, lokasi entity.Lokasi) error {
	err := r.db.Table("lokasi").Create(&lokasi).Error
	if err != nil {
		return err
	}
	return nil
}
