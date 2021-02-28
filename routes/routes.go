package routes

import (
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
			order.PUT("/:order_id", service.PutStatusOrder)
		}
	}

	return r
}
