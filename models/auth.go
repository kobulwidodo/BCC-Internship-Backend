package models

import (
	"bengkel/config"
	"bengkel/entity"
	"errors"

	"golang.org/x/crypto/bcrypt"
)


func RegitserUser(user *entity.RegisterUser) (err error) {
	if err := checkDataExist(user.Email, user.Username); err != nil {
		return err;
	}
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
	if err := config.DB.Save(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(loginUser *entity.LoginUser, user *entity.User) (err error) {
	if err := config.DB.First(&user, "email = ?", loginUser.Email).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		return err
	}
	return nil
}

func checkDataExist(email string, username string) (err error) {
	var user entity.User
	if err := config.DB.First(&user, "email = ? OR username = ?", email, username).Error;err == nil {
		return errors.New("Data sudah tersedia")
	}
	return nil
}
