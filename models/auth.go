package models

import (
	"bengkel/entity"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func RegitserUser(DB *gorm.DB, user *entity.RegisterUser) error { // change err
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost);
	if err != nil {
		return err
	}
	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Username: user.Username,
		Password: string(PasswordHash),
		Role: "Buyer",
	}
	if err := DB.Save(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(DB *gorm.DB, loginUser *entity.LoginUser, user *entity.User) (err error) {
	if err := DB.First(&user, "email = ? OR username = ?", loginUser.Email, loginUser.Email).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		return err
	}
	return nil
}

func CheckDataExist(DB *gorm.DB, email string, username string) (err error) {
	var user entity.User
	if err := DB.First(&user, "email = ? OR username = ?", email, username).Error;err == nil {
		return errors.New("Data sudah tersedia")
	}
	return nil
}
