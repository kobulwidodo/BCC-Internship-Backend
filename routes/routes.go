package routes

import (
	"bengkel/middleware"
	"bengkel/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func AddRoutes() *gin.Engine {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	
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

		product := api.Group("/product")
		{
			product.GET("/", service.GetAllProduct)
			product.POST("/new", middleware.IsAuth(), service.PostNewProduct)
			product.GET("/:product_id", service.GetProductById)
			product.PUT("/:product_id", middleware.IsAuth(), service.PutProduct)
		}

		order := api.Group("/order")
		{
			order.GET("/", service.GetAllOrder)
			order.POST("/new", service.PostNewOrder)
			order.GET("/:order_id", service.GetOrder)
			order.PUT("/:order_id", middleware.IsAuth(), service.PutStatusOrder)
		}

	}


	return r
}
