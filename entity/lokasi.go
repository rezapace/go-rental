package entity

import (
    "time"
)

type Lokasi struct {
    ID         int       `json:"id"`
    UserID     int       `json:"user_id"`
    Address    string    `json:"address"`
    City       string    `json:"city"`
    State      string    `json:"state"`
    PostalCode string    `json:"postal_code"`
    Country    string    `json:"country"`
    Latitude   float64   `json:"latitude"`
    Longitude  float64   `json:"longitude"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

// create lokasi
func CreateLokasi(id, userID int, address, city, state, postalCode, country string, latitude, longitude float64) Lokasi {
    now := time.Now()
    return Lokasi{
        ID:         id,
        UserID:     userID,
        Address:    address,
        City:       city,
        State:      state,
        PostalCode: postalCode,
        Country:    country,
        Latitude:   latitude,
        Longitude:  longitude,
        CreatedAt:  now,
        UpdatedAt:  now,
    }
}
