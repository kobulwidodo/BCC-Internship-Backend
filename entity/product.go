package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Manufacture string `json:"manufacture" binding:"required"`
	Price int `json:"price" binding:"required"`
	ImageLink string `json:"image_link" binding:"required"`
	IsAvailable bool `json:"is_available" binding:"required"`
}

type NewProduct struct {
	Name string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Manufacture string `json:"manufacture" binding:"required"`
	Price int `json:"price" binding:"required"`
	ImageLink string `json:"image_link" binding:"required"`
}
