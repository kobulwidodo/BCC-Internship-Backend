package models

import (
	"bengkel/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func PutChangePassword(DB *gorm.DB, newPassword string, user *entity.User) (err error) {
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(newPasswordHash)
	if err := DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserDetail(user *entity.User, userProfile *entity.ShowProfile) {
	userProfile.Name = user.Name
	userProfile.Email = user.Email
	userProfile.Username = user.Username
}

func CheckOldPassword(OldPassword string, password string) (err error)  {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(OldPassword)); err != nil {
		return err
	}
	return nil
}

func CheckUserLogin(DB *gorm.DB, user *entity.User, userId uint) (err error)  {
	if err := DB.First(&user, "id = ?", userId).Error; err != nil {
		return err
	}
	return nil
}
