package server

import (
	"fmt"
	"log"
	"os"

	"github.com/arvianlimansyah/moonlay-test/logger"
	// "github.com/arvianlimansyah/moonlay-test/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

var (
	app Database
)

func Run() {
	// godotenv.Load()
	logger.InitLogger()
	// routes.Run()

	fmt.Println("Test")
}

func init() {
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

	database, err := gorm.Open(postgres.Open(sqlConn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
		return
	}

	fmt.Println("++++++++++++++++++ TEST ++++++++++++++++++++")

	app.DB = database

	logger.Log("auto-migrations running...")
	// database.AutoMigrate(&models.Task{})
	// database.AutoMigrate(&models.Subtask{})
	logger.Log("auto-migration complete...")
	// database.AutoMigrateDB(app.DB)

}
