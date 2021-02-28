package models

import (
	"bengkel/config"
	"bengkel/entity"
	"math/rand"
	"time"
)


func GetAllOrders(order *[]entity.Order) (err error) {
	if err := config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func PostNewOrder(order *entity.NewOrder) (orderId string, err error)  {
	newOrderId := generateOrderID();
	dataNewOrder := entity.Order{
		OrderID: newOrderId,
		Name: order.Name,
		Email: order.Email,
		NoWhatsapp: order.NoWhatsapp,
		Order: order.Order,
		Status: "Menunggu Konfirmasi",
	}
	if err := config.DB.Save(&dataNewOrder).Error; err != nil {
		return "", err
	}
	return newOrderId, nil
}

func GetOrder(order *entity.Order, id string) (err error)  {
	if err := config.DB.First(order, "order_id", id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStatus(order *entity.Order, status string) (err error)  {
	order.Status = status
	if err := config.DB.Save(order).Error; err != nil {
		return err
	}
	return nil
}

func generateOrderID() string {
	var order entity.Order
	var b []rune
	for i := 0; i < 1; i++ {
		i--
		letters := []rune("123456789")
		rand.Seed(time.Now().UnixNano())
		b = make([]rune, 6)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		err := config.DB.First(&order, "order_id = ?", string(b)).Error
		if err != nil {
			i++
		}
	}
	return string(b)
}
