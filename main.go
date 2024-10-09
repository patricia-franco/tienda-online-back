package main

import (
	"fmt"
	"net/http"
	"tienda-online-backend/config"
	"tienda-online-backend/router"
)

func main() {
	// Conectar a la base de datos
	config.ConnectDB()

	// Inicializar las rutas
	r := router.InitializeRoutes()

	fmt.Println("Servidor Go corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
