package handler

import (
	"Rental/entity"
	"Rental/internal/http/validator"
	"Rental/internal/service"
	"Rental/common"
	"net/http"
	"log"
	"strconv"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
)

// handler
type ProdukHandler struct {
	produkService service.ProdukUsecase
}

// NewProdukHandler
func NewProdukHandler(produkService service.ProdukUsecase) *ProdukHandler {
	return &ProdukHandler{produkService}
}

// get all produk
func (h *ProdukHandler) GetAllProduk(ctx echo.Context) error {
	produk, err := h.produkService.GetAllProduk(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": produk,
	})
}

// create produk
func (h *ProdukHandler) CreateProduk(ctx echo.Context) error {
	var input struct {
		OwnerID      int     `json:"owner_id" validate:"required"`
		LocationID   int     `json:"location_id" validate:"required"`
		NameProduk   string  `json:"name_produk" validate:"required"`
		Image        string  `json:"image" validate:"required"`
		LicensePlate string  `json:"license_plate" validate:"required"`
		PricePerDay  float64 `json:"price_per_day" validate:"required"`
		Description  string  `json:"description" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	produk := entity.CreateProduk(0, input.OwnerID, input.LocationID, input.NameProduk, input.Image, input.LicensePlate, input.PricePerDay, input.Description)
	err := h.produkService.CreateProduk(ctx.Request().Context(), produk)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusCreated, "Produk created successfully")
}

// update produk
func (h *ProdukHandler) UpdateProduk(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	var input struct {
		OwnerID      int     `json:"owner_id" validate:"required"`
		LocationID   int     `json:"location_id" validate:"required"`
		NameProduk   string  `json:"name_produk" validate:"required"`
		Image        string  `json:"image" validate:"required"`
		LicensePlate string  `json:"license_plate" validate:"required"`
		PricePerDay  float64 `json:"price_per_day" validate:"required"`
		Description  string  `json:"description" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	produk := entity.UpdateProduk(id, input.OwnerID, input.LocationID, input.NameProduk, input.Image, input.LicensePlate, input.PricePerDay, input.Description)
	err = h.produkService.UpdateProduk(ctx.Request().Context(), produk)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "Produk updated successfully"})
}

// delete produk
func (h *ProdukHandler) DeleteProduk(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	err = h.produkService.DeleteProduk(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Produk deleted successfully",
	})
}

   // Get Produk by lokasi
   func (h *ProdukHandler) GetProdukByLokasi(ctx echo.Context) error {
	lokasiIDStr := ctx.Param("location_id")
	log.Printf("Received location_id: %s", lokasiIDStr) // Log untuk debug
	if lokasiIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Lokasi ID parameter is missing",
		})
	}

	lokasiID, err := strconv.Atoi(lokasiIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	produk, err := h.produkService.GetProdukByLokasi(ctx.Request().Context(), lokasiID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": produk,
	})
}

   // Get produk by price_per_day
   func (h *ProdukHandler) GetProdukByPrice(ctx echo.Context) error {
	log.Printf("Request URL: %s", ctx.Request().URL.String()) // Log for debugging
	priceStr := ctx.Param("price_per_day")
	log.Printf("Received price_per_day: %s", priceStr) // Log for debugging
	if priceStr == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Price parameter is missing",
		})
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		log.Printf("Error parsing price: %v", err) // Log for debugging
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid Price",
		})
	}

	produk, err := h.produkService.GetProdukByPrice(ctx.Request().Context(), price)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": produk,
	})
}

// search produk by name_produk
func (h *ProdukHandler) SearchProdukByName(ctx echo.Context) error {
    nameProduk := ctx.Param("name_produk")
    log.Printf("Searching for products with name containing: %s", nameProduk)
    
    produk, err := h.produkService.SearchProdukByName(ctx.Request().Context(), nameProduk)
    if err != nil {
        return ctx.JSON(http.StatusUnprocessableEntity, err)
    }
    
    log.Printf("Found products: %v", produk)
    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "data": produk,
    })
}

// get my produk with jwt
func (h *ProdukHandler) GetMyProduk(ctx echo.Context) error {
    // Get JWT token from the context
    token, ok := ctx.Get("user").(*jwt.Token)
    if !ok {
        return ctx.JSON(http.StatusUnauthorized, errors.New("missing or invalid token"))
    }

    // Extract claims from the JWT token
    claims, ok := token.Claims.(*common.JwtCustomClaims)
    if !ok {
        return ctx.JSON(http.StatusUnauthorized, errors.New("invalid token claims"))
    }

    // Log the claims for debugging
    log.Printf("Claims: %+v", claims)

    // Validate OwnerID
    if claims.OwnerID == "" {
        return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "OwnerID is required"))
    }

    // Fetch the user's products by OwnerID
    produk, err := h.produkService.GetMyProduk(ctx.Request().Context(), claims.OwnerID)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
    }

    var produkDetails []map[string]interface{}
    for _, p := range produk {
        produkDetail := map[string]interface{}{
            "id":            p.ID,
            "owner_id":      p.OwnerID,
            "location_id":   p.LocationID,
            "name_produk":   p.NameProduk,
            "image":         p.Image,
            "license_plate": p.LicensePlate,
            "price_per_day": p.PricePerDay,
            "description":   p.Description,
            "created_at":    p.CreatedAt,
            "updated_at":    p.UpdatedAt,
        }
        produkDetails = append(produkDetails, produkDetail)
    }

    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "message":        "Get all products success",
        "produk_details": produkDetails,
    })
}

// get produk by owner id
func (h *ProdukHandler) GetProdukByOwnerID(ctx echo.Context) error {
    ownerIDStr := ctx.Param("owner_id")
    log.Printf("Received owner_id: %s", ownerIDStr) // Log for debugging
    if ownerIDStr == "" {
        return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
            "error": "Owner ID parameter is missing",
        })
    }

    // Fetch the products by owner ID
    produk, err := h.produkService.GetProdukByOwnerID(ctx.Request().Context(), ownerIDStr)
    if err != nil {
        return ctx.JSON(http.StatusUnprocessableEntity, err)
    }

    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "data": produk,
    })
}

