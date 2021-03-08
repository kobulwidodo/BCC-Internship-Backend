package service

import (
	"bengkel/config"
	"bengkel/entity"
	"bengkel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func PutChangePassword(c *gin.Context) {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	var userInput entity.ChangePassword
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	var user entity.User
	var userId uint = uint(c.MustGet("jwt_user_id").(float64))
	if err := models.CheckUserLogin(DB, &user, userId); err != nil {
		c.JSON(404, gin.H{
			"status": 404,
		})
		c.Abort()
		return
	}
	if err := models.CheckOldPassword(userInput.OldPassword, user.Password); err != nil {
		c.JSON(403, gin.H{
			"status": 403,
		})
		c.Abort()
		return
	}
	if err := models.PutChangePassword(DB, userInput.NewPassword, &user); err != nil {
		c.JSON(500, gin.H{
			"status": 500,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Berhasil mengubah password!",
		"status": 200,
	})
}

func GetUserDetail(c *gin.Context)  {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	var user entity.User
	var userId uint = uint(c.MustGet("jwt_user_id").(float64))
	if err := models.CheckUserLogin(DB, &user, userId); err != nil {
		c.JSON(404, gin.H{
			"status": 404,
		})
		c.Abort()
		return
	}
	var userProfile entity.ShowProfile
	models.GetUserDetail(&user, &userProfile)
	c.JSON(200, gin.H{
		"data": userProfile,
		"status": 200,
	})
}
