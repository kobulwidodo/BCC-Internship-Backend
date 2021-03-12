package service

import (
	"bengkel/config"
	"bengkel/entity"
	"bengkel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func PostNewOrder(c *gin.Context) {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
			"status": "error",
		})
		c.Abort()
		return
	}
	var newOrder entity.NewOrder
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parameter tidak lengkap",
			"status": "error",
		})
		c.Abort()
		return
	}
	userId := int(c.MustGet("jwt_user_id").(float64))
	var cart []entity.Cart
	if err := models.GetAllCart(DB, &cart, userId); err != nil {
		c.JSON(404, gin.H{
			"message": "Cart kosong",
			"status": "error",
		})
		c.Abort()
		return
	}
	transactionId := models.GenerateTransactionId(DB)
	totalPrice, err := models.PostNewItemOrder(DB, &cart, transactionId, userId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal memasukan data (item order)",
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.PostNewOrder(DB, &newOrder, totalPrice, transactionId, userId); err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal memasukan data (order)",
			"status": "error",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"transaction_id": transactionId,
		"message": "Berhasil membuat pesanan baru",
		"status": "sukses",
	})
}

func GetAllItemOrder(c *gin.Context)  {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
			"status": "error",
		})
		c.Abort()
		return
	}
	userId := int(c.MustGet("jwt_user_id").(float64))
	var showItemOrder []entity.ShowItemOrder
	if err := models.GetAllItemOrder(DB, &showItemOrder, userId); err != nil {
		c.JSON(404, gin.H{
			"message": "Tidak dapat menemukan data",
			"status": "error",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"data": &showItemOrder,
		"message": "Sukses mendapatkan data",
		"status": "sukses",
	})
}

func GetOrderDetailById(c *gin.Context)  {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
			"status": "error",
		})
		c.Abort()
		return
	}
	transactionId := c.Param("transaction_id")
	var order entity.Order
	if err := models.GetOrderDetailById(DB, &order, transactionId); err != nil {
		c.JSON(404, gin.H{
			"message": "Tidak dapat menemukan data",
			"status": "error",
		})
		c.Abort()
		return
	}
	var showOrder entity.ShowOrder
	models.ShowOrderDetailById(&order, &showOrder)
	c.JSON(200, gin.H{
		"data": showOrder,
		"message": "Berhasil mendapatkan data",
		"status": "sukses",
	})
}
