package models

import (
	"bengkel/config"
	"bengkel/entity"
	"errors"
)


func RegitserUser(user *entity.RegisterUser) (err error) {
	if err := checkEmailExist(user.Email); err != nil {
		return err;
	}
	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	if err := config.DB.Save(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func checkEmailExist(email string) (err error) {
	var user entity.User
	if err := config.DB.First(&user, "email = ?", email).Error; err == nil {
		return errors.New("Email Sudah diAmbil")
	}
	return nil
}
