package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string
	Username string
	Password string
	Role string
}

type ChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ShowProfile struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
	Role string `json:"role"`
}
