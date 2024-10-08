package router

import (
	"github.com/gorilla/mux"
)

// InitializeRoutes exporta la función para que pueda ser utilizada fuera del paquete `router`
func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	return r
}
