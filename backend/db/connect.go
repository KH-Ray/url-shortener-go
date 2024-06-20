package db

import (
	"log"
	"url-shortener-be/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get("DATABASE_URL").(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	connStr := value
	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

    err = database.AutoMigrate(&models.ShortUrl{})

    if err != nil {
        return
    }

    DB = database
}