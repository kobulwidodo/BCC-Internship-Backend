package entity

import (
	"time"

	"gorm.io/gorm"
)


type Order struct {
	gorm.Model
	Name string
	NoHp string
	TransportationType string
	LicencePlate string
	OrderDescription string
	Complaint string
	TotalPrice int
	StnkImage string
	TransactionId string
	Status string
	UserId int
}

type NewOrder struct {
	Name string `json:"name" binding:"required"`
	NoHp string `json:"no_hp" binding:"required"`
	TransportationType string `json:"transportation_type" binding:"required"`
	LicencePlate string `json:"licence_plate" binding:"required"`
	OrderDescription string `json:"order_description" binding:"required"`
	Complaint string `json:"complaint" binding:"required"`
	StnkImage string `json:"stnk_image" binding:"required"`
}

type ShowOrder struct {
	Name string `json:"name"`
	NoHp string `json:"no_hp"`
	TransportationType string `json:"transportation_type"`
	LicencePlate string `json:"licence_plate"`
	OrderDescription string `json:"order_description"`
	Complaint string `json:"complaint"`
	CreatedAt time.Time `json:"created_at"`
	TotalPrice int `json:"total_price"`
	Status string `json:"status"`
}
