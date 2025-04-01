package database

import (
	"api/pkg/log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Error("Error loading .env file: %v", err)
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true&tls=" + dbSSLMode

	// Bağlantı işlemi
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Error creating MySQL client: %v", err)
		os.Exit(1)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Error("Error connecting to MySQL: %v", err)
		os.Exit(1)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Error("Error pinging MySQL: %v", err)
		os.Exit(1)
	}

	log.Info("Successfully connected to MySQL!")
}

func GetDB() *gorm.DB {
	return DB
}
