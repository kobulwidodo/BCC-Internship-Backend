package models

import (
	"bengkel/entity"
	"errors"

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

func GetAllCart(DB *gorm.DB, cart *[]entity.Cart, userId int) error {
	if jml := DB.Find(&cart, "user_id = ?", userId).RowsAffected; jml == 0 {
		return errors.New("Data tidak ditemukan")
	}
	return nil
}

func GetAllCartDetail(DB *gorm.DB, cart *[]entity.Cart) (entity.ShowCart) {
	var showCart []entity.GetAllCart
	price := 0
	for _, v := range *cart {
		var tempProduct entity.Product
		if err := DB.First(&tempProduct, "id = ?", v.ProductId).Error; err != nil {
			continue
		}
		temp := entity.GetAllCart{
			Id: int(v.ID),
			ProductId: v.ProductId,
			Name: tempProduct.Name,
			Description: tempProduct.Description,
			Manufacture: tempProduct.Manufacture,
			Quantity: v.Quantity,
			Price: (v.Quantity*tempProduct.Price),
			ImageLink: tempProduct.ImageLink,
		}
		price += temp.Price
		showCart = append(showCart, temp)
	}
	var tempCart entity.ShowCart
	tempCart.GetAllCart = showCart
	tempCart.TotalProduct = len(showCart)
	tempCart.TotalPrice = price
	return tempCart
}

func GetCartById(DB *gorm.DB, cart *entity.Cart, cartId int, userId int) (error) {
	if err := DB.First(&cart, "id = ? AND user_id = ?", cartId, userId).Error; err != nil {
		return err
	}
	return nil
}

func PutAddQuantity(DB *gorm.DB, cart *entity.Cart) error {
	cart.Quantity += 1
	if err := DB.Save(&cart).Error; err != nil {
		return err
	}
	return nil
}

func PutReduceQuantity(DB *gorm.DB, cart *entity.Cart) error {
	cart.Quantity -= 1
	if cart.Quantity == 0 {
		if err := DeleteCart(DB, cart); err != nil {
			return err
		}
	} else  {
		if err := DB.Save(&cart).Error; err != nil {
			return err
		}
	}
	return nil
}

func DeleteCart(DB *gorm.DB, cart *entity.Cart) (error) {
	if err := DB.Delete(&cart).Error; err != nil {
		return err
	}
	return nil
}
