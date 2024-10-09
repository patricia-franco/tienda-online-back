module tienda-online-backend

go 1.22.8

require (
	github.com/gorilla/mux v1.8.1
	gorm.io/driver/mysql v1.5.7 // Mueve mysql a la sección principal
	gorm.io/gorm v1.25.12
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.19.0 // indirect
)

replace tienda-online-backend/router => ./router
