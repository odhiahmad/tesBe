package main

import (
	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/apiuser/config"
	"github.com/odhiahmad/apiuser/controller"
	"github.com/odhiahmad/apiuser/middleware"
	"github.com/odhiahmad/apiuser/repository"
	"github.com/odhiahmad/apiuser/service"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                    = config.SetupDatabaseConnection()
	userRepository   repository.UserRepository   = repository.NewUserRepository(db)
	produkRepository repository.ProdukRepository = repository.NewProdukRepository(db)

	jwtService    service.JWTService    = service.NewJwtService()
	authService   service.AuthService   = service.NewAuthService(userRepository)
	userService   service.UserService   = service.NewUserService(userRepository)
	produkService service.ProdukService = service.NewProdukService(produkRepository)

	authController   controller.AuthController   = controller.NewAuthController(authService, jwtService)
	userController   controller.UserController   = controller.NewUserController(userService, jwtService)
	produkController controller.ProdukController = controller.NewProdukController(produkService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}
	// middleware.AuthorizeJWT(jwtService)

	userRoutes := r.Group("api/user")
	{
		userRoutes.POST("/create", userController.CreateUser)
		userRoutes.PUT("/update", userController.UpdateUser)
	}

	produkRoutes := r.Group("api/produk", middleware.AuthorizeJWT(jwtService))
	{
		produkRoutes.POST("/create", produkController.CreateProduk)
		produkRoutes.PUT("/update", produkController.UpdateProduk)
		produkRoutes.POST("/getAll", produkController.FindAll)
		produkRoutes.DELETE("/delete", produkController.Delete)
	}

	r.Run()
}
