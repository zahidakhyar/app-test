package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
	"github.com/zahidakhyar/app-test/backend/entity"
	"gorm.io/driver/postgres"
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

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:32768)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName:           "nrpostgres",
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
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
