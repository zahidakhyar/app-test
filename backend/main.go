package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zahidakhyar/app-test/backend/config"
	controller "github.com/zahidakhyar/app-test/backend/src/auth"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                 = config.SetupDatabaseConnection()
	authController controller.AuthInterface = controller.NewAuth()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := gin.Default()

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	router.Run()
}
