package router

import (
	"Rental/internal/http/handler"

	"github.com/labstack/echo/v4"
)

const (
	Admin = "Admin"
	Sewa = "Sewa"
	Penyewa = "Penyewa"
)

var (
	allRoles  = []string{Admin, Sewa, Penyewa}
	onlyAdmin = []string{Admin}
	onlySewa = []string{Sewa}
	onlyPenyewa = []string{Penyewa}
)

// membuat struct route
type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Role    []string
}

// membuat fungsi untuk mengembalikan route
// pada func ini perlu login krna private
func PublicRoutes(
	authHandler *handler.AuthHandler,
	produkHandler *handler.ProdukHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/register",
			Handler: authHandler.Registration,
		},
		{
			Method:  echo.GET,
			Path:    "/allproduk",
			Handler: produkHandler.GetAllProduk,
		},
	}
}

// membuat fungsi untuk mengembalikan route
// pada func ini tdk perlu login krna public
func PrivateRoutes(
	UserHandler *handler.UserHandler,
	produkHandler *handler.ProdukHandler,
	transaksiHandler *handler.TransaksiHandler,
	lokasiHandler *handler.LokasiHandler) []*Route {

	return []*Route{
		// Admin Routes
		{
			Method:  echo.DELETE,
			Path:    "/deleteuser/:id",
			Handler: UserHandler.DeleteUser,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/updateuser/:id",
			Handler: UserHandler.UpdateUser,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/getalluser",
			Handler: UserHandler.GetAllUser,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.POST,
			Path:    "/createuser",
			Handler: UserHandler.CreateUser,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/UpdateUserBalanceByID/:id",
			Handler: UserHandler.UpdateUserBalanceByID,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/getallproduk",
			Handler: produkHandler.GetAllProduk,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.POST,
			Path:    "/createproduk",
			Handler: produkHandler.CreateProduk,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/updateproduk/:id",
			Handler: produkHandler.UpdateProduk,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/getalltransaksi",
			Handler: transaksiHandler.GetAllTransaksi,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/gettransaksibyid/:id",
			Handler: transaksiHandler.GetTransaksiByID,
			Role:    onlyAdmin,
		},

		// All User Routes
		{
			Method:  echo.GET,
			Path:    "/users/profile",
			Handler: UserHandler.GetProfile,
			Role:    allRoles,
		},
		{
			Method:  echo.PUT,
			Path:    "/users/profile",
			Handler: UserHandler.UpdateProfile,
			Role:    allRoles,
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/deleteprofile",
			Handler: UserHandler.DeleteAccount,
			Role:    allRoles,
		},
		{
			Method: echo.GET,
			Path:   "/users/balance",
			Handler: UserHandler.GetUserBalance,
			Role:    allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/user/logout",
			Handler: UserHandler.UserLogout,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/listproduk",
			Handler: produkHandler.GetAllProduk,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/produk",
			Handler: produkHandler.GetAllProduk,
			Role:    allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/produk",
			Handler: produkHandler.CreateProduk,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.PUT,
			Path:    "/produk/:id",
			Handler: produkHandler.UpdateProduk,
			Role:    onlyAdmin,
		},
		{
			Method:  echo.GET,
			Path:    "/produk/lokasi/:location_id",
			Handler: produkHandler.GetProdukByLokasi,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/produk/price/:price_per_day",
			Handler: produkHandler.GetProdukByPrice,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/produk/name_produk/:name_produk",
			Handler: produkHandler.SearchProdukByName,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/produk/myproduk",
			Handler: produkHandler.GetMyProduk,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/produk/owner/:owner_id",
			Handler: produkHandler.GetProdukByOwnerID,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/lokasi",
			Handler: lokasiHandler.GetAllLokasi,
			Role:    allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/lokasi",
			Handler: lokasiHandler.CreateLokasi,
			Role:    allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/getransaksiprofile",
			Handler: transaksiHandler.GetTransaksiProfile,
			Role:    allRoles,
		},
	}
}

//NOTE :
//MENGAPA TERDAPAT 2 FUNC DIATAS? YAITU PUBLIC DAN PRIVATE
//KAREN DI SERVER.GO KITA BUAT GROUP API, DAN KITA MEMBAGI ROUTE YANG PERLU LOGIN DAN TIDAK PERLU LOGIN
// YAITU PUBLIC DAN PRIVATE

//note ;
//untuk menjalankan nya setelah port 8080 ditambahin /api/v1
// karna di server.go kita membuat group API
