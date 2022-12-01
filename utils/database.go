package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/arvianlimansyah/moonlay-test/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (response *gorm.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error loading .env file")
		return
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSL")
	dbTimezone := os.Getenv("DB_TIMEZONE")
	sqlConn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSslMode + " TimeZone=" + dbTimezone

	databaseConnection, errs := gorm.Open(postgres.Open(sqlConn), &gorm.Config{})
	if errs != nil {
		fmt.Println("Failed to connect to database!")
		return
	}

	Log("auto-migrations running...")
	databaseConnection.AutoMigrate(&models.Task{})
	databaseConnection.AutoMigrate(&models.Subtask{})
	Log("auto-migration complete...")

	return databaseConnection
}
