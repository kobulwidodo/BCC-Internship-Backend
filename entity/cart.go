package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId int
	ProductId int
	Quantity int
}

type PostNewCart struct {
	ProductId int `json:"product_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}
