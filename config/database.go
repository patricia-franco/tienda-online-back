package config

import (
	"fmt"
	"log"
	"tienda-online-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Variable global para almacenar la conexión a la base de datos
var DB *gorm.DB

func ConnectDB() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/tienda_online?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos: ", err)
	}

	DB = db

	//Migración  de las tablas de Product y Category
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})

	//Agregar datos iniciales si no existen las categorías
	if db.Find(&models.Category{}).RowsAffected == 0 {
		categories := []models.Category{
			{Name: "Tecnologia"},
			{Name: "Hogar"},
		}
		db.Create(&categories)
	}

	fmt.Println("Conexión a la base de datos")
}