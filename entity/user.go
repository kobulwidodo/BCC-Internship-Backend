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
