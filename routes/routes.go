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
			user.GET("/", middleware.IsAuth(), service.GetAllUser)
			user.PUT("/edit-password", middleware.IsAuth(), service.PutChangePassword)
			user.GET("/detail", middleware.IsAuth(), service.GetUserDetail)
		}

		product := api.Group("/product")
		{
			product.GET("/", service.GetAllProduct)
			product.POST("/new", middleware.IsAuth(), service.PostNewProduct)
			product.GET("/:product_id", service.GetProductById)
			product.PUT("/:product_id", middleware.IsAuth(), service.PutProduct)
			product.DELETE("/:product_id", middleware.IsAuth(), service.DeleteProduct)
		}

		cart := api.Group("/cart")
		{
			cart.GET("/", middleware.IsAuth(), service.GetAllCart)
			cart.POST("/new", middleware.IsAuth(), service.PostNewCart)
			cart.PUT("/add/:cart_id", middleware.IsAuth(), service.PutAddQuantity)
			cart.PUT("/reduce/:cart_id", middleware.IsAuth(), service.PutReduceQuantity)
			cart.DELETE("/:cart_id", middleware.IsAuth(), service.DeleteCart)
		}

		order := api.Group("/order")
		{
			order.POST("/new", middleware.IsAuth(), service.PostNewOrder)
			order.GET("/", middleware.IsAuth(), service.GetAllItemOrder)
			order.GET("/:transaction_id", service.GetOrderDetailById)
		}

	}


	return r
}
