package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID string
	Name string
	Email string
	NoWhatsapp string
	Order string
	Status string
}

type NewOrder struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	NoWhatsapp string `json:"no_whatsapp" binding:"required"`
	Order string `json:"order" binding:"required"`
}

type UpdateStatus struct {
	Status string `json:"status" binding:"required"`
}
