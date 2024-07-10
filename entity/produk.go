package entity

import (
	"time"
    "strings"

)

type Produk struct {
    ID            int       `json:"id"`
    OwnerID       int       `json:"owner_id" gorm:"column:owner_id"`
    LocationID    int       `json:"location_id" gorm:"column:location_id"`
    NameProduk    string    `json:"name_produk"`
    Image         string    `json:"image"`
    LicensePlate  string    `json:"license_plate"`
    PricePerDay   float64   `json:"price_per_day" gorm:"column:price_per_day"`
    Description   string    `json:"description"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
    DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

// create produk
func CreateProduk(id, ownerID, locationID int, nameProduk, image, licensePlate string, pricePerDay float64, description string) Produk {
    now := time.Now()
    return Produk{
        ID:           id,
        OwnerID:      ownerID,
        LocationID:   locationID,
        NameProduk:   nameProduk,
        Image:        image,
        LicensePlate: licensePlate,
        PricePerDay:  pricePerDay,
        Description:  description,
        CreatedAt:    now,
        UpdatedAt:    now,
    }
}

// update produk
func UpdateProduk(id, ownerID, locationID int, nameProduk, image, licensePlate string, pricePerDay float64, description string) Produk {
    return Produk{
        ID:           id,
        OwnerID:      ownerID,
        LocationID:   locationID,
        NameProduk:   nameProduk,
        Image:        image,
        LicensePlate: licensePlate,
        PricePerDay:  pricePerDay,
        Description:  description,
        UpdatedAt:    time.Now(),
    }
}

// delete produk
func DeleteProduk(id int, produkList []Produk) []Produk {
    now := time.Now()
    for i, produk := range produkList {
        if produk.ID == id {
            produkList[i].DeletedAt = &now
            break
        }
    }
    return produkList
}

// Get Produk by lokasi
func GetProdukByLokasi(lokasiID int, allProduk []Produk) []Produk {
    var produkList []Produk
    for _, produk := range allProduk {
        if produk.LocationID == lokasiID {
            produkList = append(produkList, produk)
        }
    }
    return produkList
}

// Get produk by price per day
func GetProdukByPricePerDay(pricePerDay float64, allProduk []Produk) []Produk {
    var produkList []Produk
    for _, produk := range allProduk {
        if produk.PricePerDay == pricePerDay {
            produkList = append(produkList, produk)
        }
    }
    return produkList
}

// search produk by name_produk
func SearchProdukByName(nameProduk string, allProduk []Produk) []Produk {
    var produkList []Produk
    for _, produk := range allProduk {
        if strings.Contains(produk.NameProduk, nameProduk) {
            produkList = append(produkList, produk)
        }
    }
    return produkList
}

// get produk by owner id
func GetProdukByOwnerID(ownerID int, allProduk []Produk) []Produk {
    var produkList []Produk
    for _, produk := range allProduk {
        if produk.OwnerID == ownerID {
            produkList = append(produkList, produk)
        }
    }
    return produkList
}

