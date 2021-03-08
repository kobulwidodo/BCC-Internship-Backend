package routes

import (
	"bengkel/middleware"
	"bengkel/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func AddRoutes() *gin.Engine {
	r := gin.Default()
	// handler := cors.Default().Handler(r)
	
	api := r.Group("/api")
	{
		api.GET("/", service.GetHome)
		
		auth := api.Group("/auth")
		{
			auth.POST("/login", service.PostLoginUser)
			auth.POST("/register", service.PostRegitserUser)
		}

		user := api.Group("/user")
		{
			user.GET("/detail", middleware.IsAuth(), service.GetUserDetail)
			user.PUT("/edit-password", middleware.IsAuth(), service.PutChangePassword)
		}

		order := api.Group("/order")
		{
			order.GET("/", service.GetAllOrder)
			order.POST("/new", service.PostNewOrder)
			order.GET("/:order_id", service.GetOrder)
			order.PUT("/:order_id", middleware.IsAuth(), service.PutStatusOrder)
		}

	}

	r.Use(cors.Default())

	return r
}
