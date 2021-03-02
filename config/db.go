package config

import (
	"bengkel/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// DB_HOST := os.Getenv("DB_HOST")
	// DB_NAME := os.Getenv("DB_NAME")
	// DB_USER := os.Getenv("DB_USER")
	// DB_PASS := os.Getenv("DB_PASS")
	dsn := "root:@tcp(127.0.0.1:3306)/gobengkel?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+")/"+DB_NAME+"?charset=utf8mb4&parseTime=True&loc=Local" // => ini error kak
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	DB.AutoMigrate(&entity.Order{})
	DB.AutoMigrate(&entity.User{})
}
