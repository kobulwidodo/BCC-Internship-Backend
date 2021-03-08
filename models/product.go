package models

import (
	"bengkel/entity"

	"gorm.io/gorm"
)

func GetAllProduct(DB *gorm.DB, product *[]entity.Product) error {
	if err := DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductById(DB *gorm.DB, product *entity.Product, id string) error  {
	if err := DB.First(product, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func PostNewProduct(DB *gorm.DB, product *entity.NewProduct) error {
	newProduct := entity.Product{
		Name: product.Name,
		Description: product.Description,
		Manufacture: product.Manufacture,
		Price: product.Price,
		ImageLink: product.ImageLink,
		IsAvailable: true,
	}
	if err := DB.Save(&newProduct).Error; err != nil {
		return err
	}
	return nil
}

func PutProduct(DB *gorm.DB, product *entity.Product, productExist *entity.Product) error  {
	productExist.Name = product.Name
	productExist.Description = product.Description
	productExist.Manufacture = product.Manufacture
	productExist.Price = product.Price
	productExist.ImageLink = product.ImageLink
	productExist.IsAvailable = product.IsAvailable
	if err := DB.Save(&productExist).Error; err != nil {
		return err
	}
	return nil
}
