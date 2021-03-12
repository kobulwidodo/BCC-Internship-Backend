package models

import (
	"bengkel/entity"
	"errors"

	"gorm.io/gorm"
)


func PostNewItemOrder(DB *gorm.DB, cart *[]entity.Cart, trxId string, userId int) (int, error) {
	var itemOrder []entity.ItemOrder
	price := 0
	for _, dataCart := range *cart {
		// fmt.Println(dataCart.ProductId)
		var tempProduct entity.Product
		if err := DB.First(&tempProduct, "id = ?", dataCart.ProductId).Error; err != nil {
			continue
		}
		temp := entity.ItemOrder{
			Name: tempProduct.Name,
			Quantity: dataCart.Quantity,
			TotalPrice: (dataCart.Quantity*tempProduct.Price),
			ImageLink: tempProduct.ImageLink,
			TransactionId: trxId,
			ProductId: int(tempProduct.ID),
			UserId: userId,
		}
		price += temp.TotalPrice
		itemOrder = append(itemOrder, temp)
		DB.Delete(&dataCart)
	}
	// panic(nil)
	if err := DB.Save(&itemOrder).Error; err != nil {
		return price, err
	}
	return price, nil
}

func GetAllItemOrder(DB *gorm.DB, showItemOrder *[]entity.ShowItemOrder, userId int) error {
	var itemOrder []entity.ItemOrder
	if jml := DB.Find(&itemOrder, "user_id = ?", userId).Scan(&showItemOrder).RowsAffected; jml == 0 {
		return errors.New("Data tidak ditemukan")
	}
	return nil
}
