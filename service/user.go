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
			"message": "Parameter tidak lengkap",
			"status": "error",
		})
		c.Abort()
		return
	}
	var user entity.User
	var userId uint = uint(c.MustGet("jwt_user_id").(float64))
	if err := models.CheckUserLogin(DB, &user, userId); err != nil {
		c.JSON(404, gin.H{
			"message": "User login tidak ditemukan",
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.CheckOldPassword(userInput.OldPassword, user.Password); err != nil {
		c.JSON(403, gin.H{
			"message": "Password lama tidak sesuai",
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.PutChangePassword(DB, userInput.NewPassword, &user); err != nil {
		c.JSON(500, gin.H{
			"message": "Terjadi kesalahan server",
			"status": "error",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Berhasil mengubah password!",
		"status": "sukses",
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
			"message": "Data user tidak ditemukan",
			"status": "error",
		})
		c.Abort()
		return
	}
	var userProfile entity.ShowProfile
	models.GetUserDetail(&user, &userProfile)
	c.JSON(200, gin.H{
		"data": userProfile,
		"message": "Berhasil mendapatkan data",
		"status": "sukses",
	})
}

func GetAllUser(c *gin.Context)  {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	var role string = string(c.MustGet("jwt_user_role").(string))
	if role != "Staff" && role != "Owner" {
		c.JSON(403, gin.H{
			"message": "Tidak memiliki akses",
			"status": "error",
		})
		c.Abort()
		return
	}
	var user []entity.ShowProfile
	if err := models.GetAllUser(DB, &user); err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal mendapatkan data",
			"status": "error",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"data": user,
		"message": "Berhasil mendapatkan seluruh data",
		"status": "sukses",
	})

}
