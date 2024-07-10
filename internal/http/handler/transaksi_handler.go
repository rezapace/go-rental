package handler

import (
	"Rental/internal/service"
	"Rental/common"
	"Rental/entity"
	"net/http"
	"strconv"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
)

type TransaksiHandler struct {
	transaksiService service.TransaksiUsecase
}

// NewTransaksiHandler
func NewTransaksiHandler(transaksiService service.TransaksiUsecase) *TransaksiHandler {
	return &TransaksiHandler{transaksiService}
}

// GetTransaksiByID
func (h *TransaksiHandler) GetTransaksiByID(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id"`
	}

	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}
	input.ID = id

	transaksi, err := h.transaksiService.GetTransaksiByID(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Unable to fetch transaction",
		})
	}

	return ctx.JSON(http.StatusOK, transaksi)
}

// GetAllTransaksi
func (h *TransaksiHandler) GetAllTransaksi(ctx echo.Context) error {
	transaksi, err := h.transaksiService.GetAllTransaksi(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Unable to fetch transactions",
		})
	}

	return ctx.JSON(http.StatusOK, transaksi)
}

// GetTransaksiProfile retrieves transactions based on the user's profile using the JWT ID
func (h *TransaksiHandler) GetTransaksiProfile(ctx echo.Context) error {
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

    // Get all transactions by user ID
    transaksi, err := h.transaksiService.GetTransaksiByUserID(ctx.Request().Context(), claims.ID)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
    }

    var transaksiDetails []map[string]interface{}
    for _, t := range transaksi {
        var produk *entity.Produk
        if t.ProdukID != nil {
            produk, err = h.transaksiService.GetProdukByID(ctx.Request().Context(), *t.ProdukID)
            if err != nil {
                return ctx.JSON(http.StatusInternalServerError, echo.NewHTTPError(http.StatusInternalServerError, err.Error()))
            }
        }

        transaksiDetail := map[string]interface{}{
            "user_id":        t.UserID,
            "produk":         produk,
            "tipe_transaksi": t.TipeTransaksi,
            "jumlah":         t.Jumlah,
            "created_at":     t.CreatedAt,
        }
        transaksiDetails = append(transaksiDetails, transaksiDetail)
    }

    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "message":           "Get all transactions success",
        "transaksi_details": transaksiDetails,
    })
}
