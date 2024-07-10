package handler

import (
	"Rental/entity"
	"Rental/internal/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// handler 
type LokasiHandler struct {
    lokasiService service.LokasiUsecase
}

// NewLokasiHandler creates a new LokasiHandler
func NewLokasiHandler(lokasiService service.LokasiUsecase) *LokasiHandler {
    return &LokasiHandler{lokasiService: lokasiService}
}

// get all lokasi
func (h *LokasiHandler) GetAllLokasi(ctx echo.Context) error {
	lokasi, err := h.lokasiService.GetAllLokasi(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": lokasi,
	})
}

// create Lokasi
func (h *LokasiHandler) CreateLokasi(ctx echo.Context) error {
	var input struct {
		UserID     int     `json:"user_id" validate:"required"`
		Address    string  `json:"address" validate:"required"`
		City       string  `json:"city" validate:"required"`
		State      string  `json:"state" validate:"required"`
		PostalCode string  `json:"postal_code" validate:"required"`
		Country    string  `json:"country" validate:"required"`
		Latitude   float64 `json:"latitude" validate:"required"`
		Longitude  float64 `json:"longitude" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid input",
		})
	}

	// Manually validate required fields
	if input.UserID == 0 || input.Address == "" || input.City == "" || input.State == "" || 
	   input.PostalCode == "" || input.Country == "" || input.Latitude == 0 || input.Longitude == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "All fields are required",
		})
	}

	lokasi := entity.Lokasi{
		UserID:     input.UserID,
		Address:    input.Address,
		City:       input.City,
		State:      input.State,
		PostalCode: input.PostalCode,
		Country:    input.Country,
		Latitude:   input.Latitude,
		Longitude:  input.Longitude,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Save the lokasi
	err := h.lokasiService.CreateLokasi(ctx.Request().Context(), lokasi)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": lokasi,
	})
}