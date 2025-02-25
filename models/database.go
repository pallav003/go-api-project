package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=aws-0-ap-south-1.pooler.supabase.com user=postgres.qsfbrryskxsynsrbcrkf password=[Your-password] dbname=postgres port=5432 sslmode=require"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&User{}) // Auto-migrate user table
	DB = database
}
