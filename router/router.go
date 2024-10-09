package router

import (
	"github.com/gin-gonic/gin"
	"tienda-online-backend/api"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()

	//Rutas CRUD para productos
	router.POST("/api/products", api.CreateProduct)
	router.GET("/api/products", api.GetProducts)
	router.GET("/api/products/:id", api.GetProductByID)
	router.PUT("/api/products/:id", api.UpdateProduct)
	router.DELETE("/api/products/:id", api.DeleteProduct)

	return router
}