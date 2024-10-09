package api

import (
	"net/http"
	"tienda-online-backend/config"
	"tienda-online-backend/models"
	"github.com/gin-gonic/gin"
)

//Crear nuevo producto
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}

//Obtener un producto por ID
func GetProductByID(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := config.DB.Preload("Category").Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, product)
}

//Obtener todos los productos
func GetProducts(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.Preload("Category").Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, product)
}

//Actualizar un producto
func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

//Eliminar un producto
func DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado con exito"})
}

