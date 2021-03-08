package config

import (
	"bengkel/entity"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	// dsn := "root:@tcp(127.0.0.1:3306)/gobengkel?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+")/"+DB_NAME+"?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	// DB.AutoMigrate(&entity.Order{})
	if err := DB.AutoMigrate(&entity.User{}); err != nil {
		return nil, err
	}
	return DB, nil
}
