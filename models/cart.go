package models

import (
	"bengkel/entity"

	"gorm.io/gorm"
)


func PostNewCart(DB *gorm.DB, cart *entity.PostNewCart, userId int) error {
	newCart := entity.Cart{
		UserId: userId,
		ProductId: cart.ProductId,
		Quantity: cart.Quantity,
	}
	if err := DB.Save(&newCart).Error; err != nil {
		return err
	}
	return nil
}
