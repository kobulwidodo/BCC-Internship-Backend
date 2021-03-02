package service

import (
	"bengkel/entity"
	"bengkel/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAllOrder(c *gin.Context) {
	var order []entity.Order
	err := models.GetAllOrders(&order)
	if err != nil {
		c.JSON(404, gin.H{
			"status": http.StatusNotFound,
		})
	} else {
		c.JSON(400, gin.H{
			"data": order,
			"status": http.StatusOK,
		})
	}
}

func PostNewOrder(c *gin.Context)  {
	var order entity.NewOrder
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	newOrderId, err := models.PostNewOrder(&order)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(404, gin.H{
			"status": http.StatusNotFound,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Berhasil membuat pesanan baru",
		"order_id": newOrderId,
		"status": http.StatusOK,
	})
}

func GetOrder(c *gin.Context)  {
	var order entity.Order
	order_id := c.Param("order_id")
	err := models.GetOrder(&order, order_id)
	if err != nil {
		c.JSON(404, gin.H{
			"status": http.StatusNotFound,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"data": order,
		"status": http.StatusOK,
	})
}

func PutStatusOrder(c *gin.Context)  {
	var order entity.Order
	orderId := c.Param("order_id")
	err := models.GetOrder(&order, orderId)
	if err != nil {
		c.JSON(404, gin.H{
			"status": http.StatusNotFound,
		})
		c.Abort()
		return
	}
	var isAdmin string = c.MustGet("jwt_user_role").(string)
	if isAdmin != "Admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
		})
		c.Abort()
		return
	}
	var newStatus entity.UpdateStatus
	if err := c.BindJSON(&newStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	err = models.UpdateStatus(&order, newStatus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Berhasil Update Status",
		"status": http.StatusOK,
	})
}
