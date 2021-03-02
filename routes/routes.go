package routes

import (
	"bengkel/middleware"
	"bengkel/service"

	"github.com/gin-gonic/gin"
)


func AddRoutes() *gin.Engine {
	r := gin.Default()
	
	api := r.Group("/api")
	{
		api.GET("/", service.GetHome)

		order := api.Group("/order")
		{
			order.GET("/", service.GetAllOrder)
			order.POST("/new", service.PostNewOrder)
			order.GET("/:order_id", service.GetOrder)
			order.PUT("/:order_id", middleware.IsAuth(), service.PutStatusOrder)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/login", service.PostLoginUser)
			auth.POST("/register", service.PostRegitserUser)
		}
	}

	return r
}
