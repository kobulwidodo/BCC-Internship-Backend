package service

import (
	"bengkel/config"
	"bengkel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func PostNewAdmin(c *gin.Context) {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.CheckDataAdmin(DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Akun admin sudah tersedia",
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.PostNewAdmin(DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal memasukan data",
			"status": "error",
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "sukses migrate user admin",
		"status": "sukses",
	})
}
