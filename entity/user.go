package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string
	NoHp string
	Password string
	Role string
}

type ChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
