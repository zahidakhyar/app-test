package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/zahidakhyar/app-test/backend/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	env := godotenv.Load()

	if env != nil {
		panic("Error loading .env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.User{})

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("failed to close database connection")
	}

	dbSQL.Close()
}
