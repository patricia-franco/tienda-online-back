package main

import (
	"fmt"
	"net/http"
	"tienda-online-backend/router" // Importa el paquete 'router' de la carpeta 'back/router'
)

func main() {
	r := router.InitializeRoutes() // Inicializa las rutas desde el enrutador
	fmt.Println("Servidor Go corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", r) // Inicia el servidor en el puerto 8080
}
