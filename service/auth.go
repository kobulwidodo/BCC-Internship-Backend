package service

import (
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
	var userInput entity.RegisterUser
	if err := c.BindJSON(&userInput); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		c.Abort()
		return
	}
	if err := models.CheckDataExist(userInput.Email, userInput.NoHp); err != nil {
		c.JSON(500, gin.H{
			"status": 500,
		})
		c.Abort()
		return
	}
	if err := models.RegitserUser(&userInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		c.Abort()
		return
	}
	c.JSON(201, gin.H{
		"message" : "Berhasil membuat Akun",
		"status": 201,
	})
}

func PostLoginUser(c *gin.Context)  {
	var loginUser entity.LoginUser
	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		c.Abort()
		return
	}

	var user entity.User
	err := models.LoginUser(&loginUser, &user)
	if err != nil {
		c.JSON(404, gin.H{
			"status": 404,
		})
		c.Abort()
		return
	}
	var jwtToken = generateToken(&user)

	c.JSON(200, gin.H{
		"status": 200,
		"message": "Berhasil Login!",
		"token": jwtToken,
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
