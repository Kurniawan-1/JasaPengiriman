package connection

import (
	"fmt"
	"jasaPengiriman/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal("Can't Load Env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't Connect To Database")
	}

	db.AutoMigrate(&models.Pelanggan{}, &models.Kurir{}, &models.Barang{}, &models.Pengiriman{})

	return db
}
