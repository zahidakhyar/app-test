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

	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
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

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("app-test [Backend]"),
		newrelic.ConfigLicense("90b52b3201fa78c293523de43cfb61364a22NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigEnabled(true),
	)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")
	router.Use(
		cors.New(config),
		nrgin.Middleware(app),
	)

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	profileRoutes := router.Group("api/auth/profile", middleware.AuthorizeJwt(jwtService))
	{
		profileRoutes.PUT("", authController.Update)
		profileRoutes.GET("", authController.Profile)
	}

	router.Run()
}
