package models

import (
	"bengkel/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func PostNewAdmin(DB *gorm.DB) error {
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte("admin010120"), bcrypt.MinCost);
	if err != nil {
		return err
	}
	user := entity.User{
		Name: "Admin",
		Email: "fadhilhan01@gmail.com",
		Username: "adminganteng",
		Password: string(PasswordHash),
		Role: "Owner",
	}
	if err := DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func CheckDataAdmin(DB *gorm.DB) error {
	var user entity.User
	if err := DB.First(&user, "Role = ?", "Owner").Error; err == nil {
		return err
	}
	return nil
}
