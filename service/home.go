package service

import "github.com/gin-gonic/gin"

func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to home page",
	});
}
