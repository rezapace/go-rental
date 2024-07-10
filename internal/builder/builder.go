package builder

import (
	"Rental/internal/config"
	"Rental/internal/http/handler"
	"Rental/internal/http/router"
	"Rental/internal/repository"
	"Rental/internal/service"

	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

// ini untuk public
func BuildPublicRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	registrationRepository := repository.NewRegistrationRepository(db)
	registrationService := service.NewRegistrationService(registrationRepository)

	// user
	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)

	// produk
	produkRepository := repository.NewProdukRepository(db)
	produkService := service.NewProdukService(produkRepository)
	produkHandler := handler.NewProdukHandler(produkService)

	// auth
	authHandler := handler.NewAuthHandler(registrationService, loginService, tokenService)

	// Update the line below with the additional TicketHandler argument
	return router.PublicRoutes(authHandler, produkHandler)
}

// ini untuk private 
func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client) []*router.Route {
	// Create a user handler
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// produk
	produkRepository := repository.NewProdukRepository(db)
	produkService := service.NewProdukService(produkRepository)
	produkHandler := handler.NewProdukHandler(produkService)

	// transaksi
	transaksiRepository := repository.NewTransaksiRepository(db)
	transaksiService := service.NewTransaksiService(transaksiRepository)
	transaksiHandler := handler.NewTransaksiHandler(transaksiService)

	// lokasi
    lokasiRepository := repository.NewLokasiRepository(db)
    lokasiService := service.NewLokasiService(lokasiRepository)
    lokasiHandler := handler.NewLokasiHandler(lokasiService)

	// Menggunakan PrivateRoutes dengan kedua handler
	return router.PrivateRoutes(userHandler, produkHandler, transaksiHandler, lokasiHandler)
}
