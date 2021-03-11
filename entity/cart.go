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

type GetAllCart struct {
	Id int
	ProductId int
	Name string
	Description string
	Manufacture string
	Quantity int
	Price int
	ImageLink string
}

type ShowCart struct {
	GetAllCart interface{} `json:"cart"`
	TotalProduct interface{} `json:"total_product"`
	TotalPrice interface{} `json:"total_price"`
}
