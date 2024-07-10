package handler

//NOTE :
// FOLDER INI UNTUK MEMANGGIL SERVICE DAN REPOSITORY
import (
	"Rental/common"
	"Rental/entity"
	"Rental/internal/http/validator"
	"Rental/internal/service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserUsecase
}

// melakukan instace dari user handler
func NewUserHandler(userService service.UserUsecase) *UserHandler {
	return &UserHandler{userService}
}

// func untuk melakukan getAll User
func (h *UserHandler) GetAllUser(ctx echo.Context) error {
	users, err := h.userService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
}

// func untuk melakukan createUser update versi reza v5 halo
func (h *UserHandler) CreateUser(ctx echo.Context) error {
	var input struct {
		Name     string  `json:"name" validate:"required"`
		Email    string  `json:"email" validate:"email"`
		Number   string  `json:"number" validate:"min=11,max=13"`
		Role     string  `json:"role" validate:"oneof=Admin Sewa Penyewa"`
		Password string  `json:"password"`
		Balance  float64 `json:"balance"`
	}
	//ini func untuk error checking
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}
	user := entity.NewUser(input.Name, input.Email, input.Number, input.Password, input.Role, input.Balance)
	err := h.userService.CreateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}
	//kalau retrun nya kaya gini akan tampil pesan "User created successfully"
	return ctx.JSON(http.StatusCreated, "User created successfully")
}

// func untuk melakukan updateUser by id
func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	var input struct {
		ID       int64  `param:"id" validate:"required"`
		Name     string `json:"name"`
		Email    string `json:"email" validate:"email"`
		Number   string `json:"number" validate:"min=11,max=13"`
		Role     string `json:"role" validate:"oneof=Admin Buyer"`
		Password string `json:"password"`
		Balance  float64 `json:"balance"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user := entity.UpdateUser(input.ID, input.Name, input.Email, input.Number, input.Role, input.Password, input.Balance)

	err := h.userService.UpdateUser(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "succesfully update user"})
}

// func untuk melakukan getUser by id
func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		// Jika tidak dapat mengonversi ID menjadi int64, kembalikan respons error
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}
	user, err := h.userService.GetUserByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":       user.ID,
			"name":     user.Name,
			"email":    user.Email,
			"number":   user.Number,
			"password": user.Password,
			"created":  user.CreatedAt,
			"updated":  user.UpdatedAt,
		},
	})
}

// DeleteUser func untuk melakukan delete user by id
func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.userService.Delete(ctx.Request().Context(), input.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

// Update User Self
func (h *UserHandler) UpdateProfile(ctx echo.Context) error {
	var input struct {
		ID       int64
		Name     string `json:"name"`
		Email    string `json:"email" validate:"email"`
		Number   string `json:"number" validate:"min=11,max=13"`
		Password string `json:"password"`
		Balance  float64 `json:"balance"`
	}

	// Mengambil nilai 'claims' dari JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	// Mengisi ID dari klaim ke input
	input.ID = claimsData.ID

	// Update user
	user := entity.UpdateProfile(input.ID, input.Name, input.Email, input.Number, input.Password)

	// Memanggil service untuk update user
	err := h.userService.UpdateProfile(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully update user"})
}

// get profile
func (h *UserHandler) GetProfile(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	// Fetch user profile using the user ID
	user, err := h.userService.GetProfile(ctx.Request().Context(), claimsData.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Return the user profile
	return ctx.JSON(http.StatusOK, user)
}

// Get user balance
func (h *UserHandler) GetUserBalance(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	// Fetch user balance using the user ID
	balance, err := h.userService.GetUserBalance(ctx.Request().Context(), claimsData.ID)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Return the user balance
	return ctx.JSON(http.StatusOK, balance.Balance)
}

// delete account
func (h *UserHandler) DeleteAccount(ctx echo.Context) error {
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	// Menggunakan ID dari klaim JWT
	idToDelete := claimsData.ID

	err := h.userService.Delete(ctx.Request().Context(), idToDelete)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

// upgrade balance
func (h *UserHandler) UpgradeBalance(ctx echo.Context) error {
	// Retrieve user ID from JWT claims
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	userID := claimsData.ID

	// Fetch current balance for the user
	currentUser, err := h.userService.GetUserByID(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "unable to fetch user information")
	}

	// Extract input data
	var input struct {
		Balance float64 `json:"balance"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Add the new balance to the current balance
	newBalance := currentUser.Balance + input.Balance

	// Update user balance
	currentUser.Balance = newBalance
	err = h.userService.UpgradeBalance(ctx.Request().Context(), currentUser)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully updated user balance"})
}

// logout
func (h *UserHandler) UserLogout(ctx echo.Context) error {
	// Retrieve user claims from the JWT token
	claims, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user claims")
	}

	// Extract user information from claims
	claimsData, ok := claims.Claims.(*common.JwtCustomClaims)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, "unable to get user information from claims")
	}

	userID := claimsData.ID

	// Create a *entity.User instance with the userID
	user := &entity.User{ID: userID}

	// Invalidate the JWT token
	err := h.userService.UserLogout(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "unable to invalidate JWT token")
	}

	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully logged out"})
}

// updateuserbalance by id
func (h *UserHandler) UpdateUserBalanceByID(ctx echo.Context) error {
    // Extract user ID from the URL parameter
    idStr := ctx.Param("id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
            "error": "Invalid ID",
        })
    }

    // Extract input data
    var input struct {
        Balance float64 `json:"balance" validate:"required"`
    }

    if err := ctx.Bind(&input); err != nil {
        return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
    }

    // Fetch current user by ID
    currentUser, err := h.userService.GetUserByID(ctx.Request().Context(), id)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, "unable to fetch user information")
    }

    // Update user balance
    currentUser.Balance = input.Balance
    err = h.userService.UpdateUserBalanceByID(ctx.Request().Context(), currentUser)
    if err != nil {
        return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
    }

	return ctx.JSON(http.StatusOK, map[string]string{"success": "successfully updated user balance"})
}
