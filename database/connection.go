package database

import (
	"EFB-User-Backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "efbusercloudrun:Efbusercloudrun123$@tcp(159.223.90.91:3306)/efb_users?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.Airline{}, &models.User{}, &models.UserSession{})
	DB = database

	log.Println("✅ Database connected successfully.")
}
