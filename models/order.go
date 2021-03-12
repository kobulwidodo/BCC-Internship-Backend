package models

import (
	"bengkel/entity"
	"math/rand"
	"time"

	"gorm.io/gorm"
)


func PostNewOrder(DB *gorm.DB, newOrder *entity.NewOrder, totalPrice int, trxId string, userId int) error {
	order := entity.Order{
		Name: newOrder.Name,
		NoHp: newOrder.NoHp,
		TransportationType: newOrder.TransportationType,
		LicencePlate: newOrder.LicencePlate,
		OrderDescription: newOrder.OrderDescription,
		Complaint: newOrder.Complaint,
		TotalPrice: totalPrice,
		StnkImage: newOrder.StnkImage,
		TransactionId: trxId,
		Status: "Menunggu Konfirmasi",
		UserId: userId,
	}
	if err := DB.Save(&order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderDetailById(DB *gorm.DB, order *entity.Order, transactionId string) error {
	if err := DB.First(&order, "transaction_id = ?", transactionId).Error; err != nil {
		return err
	}
	return nil
}

func ShowOrderDetailById(order *entity.Order, showOrder *entity.ShowOrder)  {
	showOrder.Name = order.Name
	showOrder.NoHp = order.NoHp
	showOrder.TransportationType = order.TransportationType
	showOrder.LicencePlate = order.LicencePlate
	showOrder.OrderDescription = order.OrderDescription
	showOrder.Complaint = order.Complaint
	showOrder.CreatedAt = order.CreatedAt
	showOrder.TotalPrice = order.TotalPrice
	showOrder.Status = order.Status
}

func GenerateTransactionId(DB *gorm.DB) string {
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
		err := DB.First(&order, "transaction_id = ?", string(b)).Error
		if err != nil {
			i++
		}
	}
	return string(b)
}
