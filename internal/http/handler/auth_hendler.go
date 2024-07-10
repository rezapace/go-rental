package handler

import (
	"Rental/entity"
	"Rental/internal/http/validator"
	"Rental/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	registrationService service.RegistrationUseCase // untuk regist
	loginService        service.LoginUseCase        //untuk memanggil service yang ngelakuin pengecekan user.
	tokenService        service.TokenUsecase        //untuk memanggil func akses token
}

// ini func untuk type AuthHandler
func NewAuthHandler(
	registartionService service.RegistrationUseCase,
	loginService service.LoginUseCase,
	tokenService service.TokenUsecase,

) *AuthHandler {
	return &AuthHandler{
		registrationService: registartionService,
		loginService:        loginService,
		tokenService:        tokenService,
	}
}

// func ini untuk login
func (h *AuthHandler) Login(ctx echo.Context) error {
	// pengecekan request
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	if err := ctx.Bind(&input); err != nil { // di cek pake validate buat masukin input
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// untuk manggil login service di folder service
	user, err := h.loginService.Login(ctx.Request().Context(), input.Email, input.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	// untuk manggil token service di folder service
	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	data := map[string]interface{}{
		"access_token": accessToken,
	}
	return ctx.JSON(http.StatusOK, data)
}

// Public Register
func (h *AuthHandler) Registration(ctx echo.Context) error {
	// pengecekan request
	var input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Roles    string `json:"roles" validate:"required,oneof=Admin Sewa Penyewa"`
		Number   string `json:"number" validate:"required,min=11,max=13"`
	}

	if err := ctx.Bind(&input); err != nil { // di cek pake validate buat masukin input
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// untuk manggil registration service di folder service
	user := entity.Register(input.Email, input.Password, input.Roles, input.Number)
	err := h.registrationService.Registration(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":      "User registration successfully",
		"access_token": accessToken,
	})
}


