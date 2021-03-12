package entity

import (
	"gorm.io/gorm"
)


type ItemOrder struct {
	gorm.Model
	Name       string
	Quantity   int
	TotalPrice int
	ImageLink  string
	TransactionId string
	ProductId  int
	UserId     int
}

type ShowItemOrder struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	TotalPrice int `json:"total_price"`
	ImageLink string `json:"image_link"`
	CreatedAt string `json:"created_at"`
	TransactionId string `json:"transaction_id"`
}
