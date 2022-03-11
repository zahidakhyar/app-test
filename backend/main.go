package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/zahidakhyar/app-test/backend/config"
	"github.com/zahidakhyar/app-test/backend/middleware"
	controller "github.com/zahidakhyar/app-test/backend/src/auth"
	auth_service "github.com/zahidakhyar/app-test/backend/src/auth/service"
	user_service "github.com/zahidakhyar/app-test/backend/src/user/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                           = config.SetupDatabaseConnection()
	userService    user_service.UserServiceInterface  = user_service.NewUserService(db)
	jwtService     auth_service.JwtServiceInterface   = auth_service.NewJwtService()
	authService    auth_service.AuthServiceInterface  = auth_service.NewAuthService(userService)
	authController controller.AuthControllerInterface = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := gin.Default()

	router.Use(cors.Default())

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	profileRoutes := router.Group("api/auth/profile", middleware.AuthorizeJwt(jwtService))
	{
		profileRoutes.PUT("/", authController.Update)
		profileRoutes.GET("/", authController.Profile)
	}

	router.Run()
}
