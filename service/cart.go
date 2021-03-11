package service

import (
	"bengkel/config"
	"bengkel/entity"
	"bengkel/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func PostNewCart(c *gin.Context) {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	var newCart entity.PostNewCart
	if err := c.Bind(&newCart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parameter tidak lengkap",
			"status": "error",
		})
		c.Abort()
		return
	}
	var product entity.Product
	productId := strconv.Itoa(newCart.ProductId)
	if err := models.GetProductById(DB, &product, productId); err != nil {
		c.JSON(404, gin.H{
			"message": "Product tidak ditemukan",
			"status": "error",
		})
		c.Abort()
		return
	}
	userId := int(c.MustGet("jwt_user_id").(float64))
	if err := models.PostNewCart(DB, &newCart, userId); err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal memasukan data",
			"status": "sukses",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Berhasil memasukan data",
		"status": "error",
	})
}
