package entity

import (
	"time"
)

type Transaksi struct {
	ID        	  int64  `gorm:"primaryKey"`
	UserID        int64     `json:"user_id"`
	ProdukID      *int64    `json:"produk_id,omitempty"` // Use pointer to allow NULL values
	TipeTransaksi string    `json:"tipe_transaksi"`
	Jumlah        float64   `json:"jumlah"`
	CreatedAt     time.Time `json:"created_at"`
}

// NewTransaksi creates a new Transaksi instance
func NewTransaksi(userID int64, produkID *int64, tipeTransaksi string, jumlah float64) *Transaksi {
	return &Transaksi{
		UserID:        userID,
		ProdukID:      produkID,
		TipeTransaksi: tipeTransaksi,
		Jumlah:        jumlah,
		CreatedAt:     time.Now(),
	}
}

func (Transaksi) TableName() string {
	return "transaksi"
}