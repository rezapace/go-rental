package repository

import (
	"context"
	"errors"
	"time"
	"log"

	"Rental/entity"

	"gorm.io/gorm"
)

// repository
type ProdukRepository struct {
	db *gorm.DB
}

// NewProdukRepository
func NewProdukRepository(db *gorm.DB) *ProdukRepository {
	return &ProdukRepository{db: db}
}

// get all produk
func (r *ProdukRepository) GetProduk(ctx context.Context) ([]entity.Produk, error) {
	var produk []entity.Produk
	err := r.db.Table("produk").Find(&produk).Error
	if err != nil {
		return nil, err
	}
	return produk, nil
}

// create produk
func (r *ProdukRepository) CreateProduk(ctx context.Context, produk entity.Produk) error {
	err := r.db.Table("produk").Create(&produk).Error
	if err != nil {
		return err
	}
	return nil
}

// update produk
func (r *ProdukRepository) UpdateProduk(ctx context.Context, produk entity.Produk) error {
	err := r.db.Table("produk").Save(&produk).Error
	if err != nil {
		return err
	}
	return nil
}

// delete produk
func (r *ProdukRepository) DeleteProduk(ctx context.Context, id int) error {
	var produk entity.Produk
	err := r.db.Table("produk").First(&produk, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	now := time.Now()
	produk.DeletedAt = &now
	err = r.db.Table("produk").Save(&produk).Error
	if err != nil {
		return err
	}
	return nil
}

   // Get Produk by lokasi
   func (r *ProdukRepository) GetProdukByLokasi(ctx context.Context, lokasiID int) ([]entity.Produk, error) {
	var produk []entity.Produk
	err := r.db.Table("produk").Where("location_id = ?", lokasiID).Find(&produk).Error  // Gunakan "location_id"
	if err != nil {
		return nil, err
	}
	return produk, nil
}

   // Get produk by price
   func (r *ProdukRepository) GetProdukByPrice(ctx context.Context, price float64) ([]entity.Produk, error) {
	var produk []entity.Produk
	err := r.db.Table("produk").Where("price_per_day = ?", price).Find(&produk).Error
	if err != nil {
		return nil, err
	}
	return produk, nil
}

// search produk by name_produk
func (r *ProdukRepository) SearchProdukByName(ctx context.Context, nameProduk string) ([]entity.Produk, error) {
	var produk []entity.Produk
	query := r.db.Table("produk").Where("name_produk ILIKE ?", "%"+nameProduk+"%")
	err := query.Find(&produk).Error
	if err != nil {
		return nil, err
	}
	log.Printf("Executed query: %v", query.Debug().Statement.SQL.String())
	return produk, nil
}

// get my produk with jwt
func (r *ProdukRepository) GetMyProduk(ctx context.Context, OwnerID string) ([]entity.Produk, error) {
	var produk []entity.Produk
	err := r.db.Table("produk").Where("owner_id = ?", OwnerID).Find(&produk).Error
	if err != nil {
		return nil, err
	}
	return produk, nil
}

// get produk by owner id
func (r *ProdukRepository) GetProdukByOwnerID(ctx context.Context, ownerID string) ([]entity.Produk, error) {
	var produk []entity.Produk
	err := r.db.Table("produk").Where("owner_id = ?", ownerID).Find(&produk).Error
	if err != nil {
		return nil, err
	}
	return produk, nil
}

