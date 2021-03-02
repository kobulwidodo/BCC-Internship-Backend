package service

import (
	"bengkel/entity"
	"bengkel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func PostRegitserUser(c *gin.Context) {
	var user entity.RegisterUser
	if err := c.Bind(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	if err := models.RegitserUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message" : "Berhasil membuat Akun",
		"status": 200,
	})
}
