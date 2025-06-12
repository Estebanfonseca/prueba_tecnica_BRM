package router

import (
	"api_users/api/handler"
	"api_users/api/middleware"
	"api_users/api/repository"
	"api_users/api/service"
	"database/sql"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "api_users/docs"
)

// NewRouter crea un nuevo enrutador con las rutas y middleware configurados.
func NewRouter(db *sql.DB) *mux.Router {
	secretKey := "my_secret_key"
	tokenDuration := 30 * time.Minute
	
	repo := repository.NewPostgresRepository(db)
	authService := service.NewAuthService(repo, secretKey, tokenDuration)
	authHandler := handler.NewAuthHandler(authService)
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	r := mux.NewRouter()
	api := r.PathPrefix("/").Subrouter()
	api.Use(middleware.AuthMiddleware(secretKey))
	api.HandleFunc("/users", userHandler.Create).Methods("POST")
	api.HandleFunc("/users", userHandler.List).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.GetByID).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	api.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	r.PathPrefix("/").Handler(httpSwagger.WrapHandler)

	return r
	
}