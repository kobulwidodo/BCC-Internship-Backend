package service

import (
	"bengkel/config"
	"bengkel/entity"
	"bengkel/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func PostRegitserUser(c *gin.Context) {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	var userInput entity.RegisterUser
	if err := c.BindJSON(&userInput); err != nil{
		c.JSON(400, gin.H{
			"message": "Parameter tidak lengkap",
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.CheckDataExist(DB, userInput.Email, userInput.Username); err != nil {
		c.JSON(500, gin.H{
			"message": "Username atau Email sudah digunakan",
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.RegitserUser(DB, &userInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message" : "Berhasil membuat Akun",
		"status": "sukses",
	})
}

func PostLoginUser(c *gin.Context)  {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	var loginUser entity.LoginUser
	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(400, gin.H{
			"message": "Parameter tidak lengkap",
			"status": "error",
		})
		c.Abort()
		return
	}

	var user entity.User
	err = models.LoginUser(DB ,&loginUser, &user)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Username atau Email tidak cocok",
			"status": "error",
		})
		c.Abort()
		return
	}
	var jwtToken = generateToken(&user)

	c.JSON(200, gin.H{
		"token": jwtToken,
		"message": "Berhasil Login!",
		"status": "sukses",
	})

}

func generateToken(user *entity.User) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"user_role": user.Role,
		"exp": time.Now().AddDate(0, 0, 7).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(JWT_SECRET))
	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}
