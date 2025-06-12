package main

import (
	"api_users/api/router"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
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