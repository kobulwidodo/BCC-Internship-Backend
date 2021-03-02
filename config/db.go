package config

import (
	"bengkel/entity"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp(127.0.0.1:3306)/"+os.Getenv("DB_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	DB.AutoMigrate(&entity.Order{})
	DB.AutoMigrate(&entity.User{})
}
