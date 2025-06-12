package main

import (
	"api_users/api/router"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	_ "github.com/swaggo/http-swagger"
)

// @title API de Usuarios
// @version 1.0
// @description Esta es una API de ejemplo para gestionar usuarios.
// @termsOfService http://swagger.io/terms/
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
func main() {
	// main es el punto de entrada de la aplicacion.
	connStr := "postgres://user:password@db:5432/mydb?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	router := router.NewRouter(db)
	log.Fatal(http.ListenAndServe(":8000", router))
}