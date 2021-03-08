package service

import (
	"bengkel/config"
	"bengkel/entity"
	"bengkel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAllProduct(c *gin.Context) {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	var product []entity.Product
	if err := models.GetAllProduct(DB, &product); err != nil {
		c.JSON(404, gin.H{
			"message": "Gagal mendapatkan data product",
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"data": product,
		"message": "Berhasil mendapatkan semua product",
		"status": "sukses",
	})
}

func PostNewProduct(c *gin.Context) {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	var newProduct entity.NewProduct
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(400, gin.H{
			"message": "Parameter tidak lengkap",
			"status": "error",
		})
		c.Abort()
		return
	}
	var role string = c.MustGet("jwt_user_role").(string)
	if role != "Staff" && role != "Owner" {
		c.JSON(403, gin.H{
			"message": "Tidak memiliki akses",
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.PostNewProduct(DB, &newProduct); err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal Menambah data Baru",
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Berhasil Menambah data baru",
		"status": "sukses",
	})
}

func GetProductById(c *gin.Context)  {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	product_id := c.Param("product_id")
	var product entity.Product
	if err := models.GetProductById(DB, &product, product_id); err != nil {
		c.JSON(404, gin.H{
			"message": "Data tidak ditemukan",
			"status": "error",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"data": product,
		"message": "Berhasil mengambil 1 data",
		"status": "sukses",
	})
}

func PutProduct(c *gin.Context)  {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	var product entity.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parameter tidak lengkap",
			"status": "error",
		})
		c.Abort()
		return
	}
	var role string = c.MustGet("jwt_user_role").(string)
	if role != "Staff" && role != "Owner" {
		c.JSON(403, gin.H{
			"message": "Tidak memiliki akses",
			"status": "error",
		})
		c.Abort()
		return
	}
	product_id := c.Param("product_id")
	var productExist entity.Product
	if err := models.GetProductById(DB, &productExist, product_id); err != nil {
		c.JSON(404, gin.H{
			"message": "Product Tidak ditemukan",
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.PutProduct(DB, &product, &productExist); err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"status": "error",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Berhasil update product",
		"status": "sukses",
	})
}

func DeleteProduct(c *gin.Context)  {
	DB, err := config.InitDB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error(),
		})
		c.Abort()
		return
	}
	product_id := c.Param("product_id")
	var product entity.Product
	if err := models.GetProductById(DB, &product, product_id); err != nil {
		c.JSON(404, gin.H{
			"message": "Product tidak ditersedia",
			"status": "error",
		})
		c.Abort()
		return
	}
	role := c.MustGet("jwt_user_role")
	if role != "Staff" && role != "Owner" {
		c.JSON(403, gin.H{
			"message": "Tidak memiliki akses",
			"status": "error",
		})
		c.Abort()
		return
	}
	if err := models.DeleteProduct(DB, &product); err != nil {
		c.JSON(500, gin.H{
			"message": "Gagal menghapus product",
			"status": "error",
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"message": "Berhasil menghapus product",
		"status": "sukses",
	})
}
