package handler

import (
	"api_users/api/middleware"
	"api_users/api/models"
	"api_users/api/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// UserHandler maneja las solicitudes de los usuarios.

type UserHandler struct {
	userService *service.UserService
}


func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Create maneja la solicitud para crear un nuevo usuario.
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.userService.Create(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// List maneja la solicitud para obtener una lista de todos los usuarios.
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// GetByID maneja la solicitud para obtener un usuario por su ID.
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID := r.Context().Value(middleware.UserIDKey).( string )

	if id != userID {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	user, err := h.userService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Update maneja la solicitud para actualizar un usuario existente.
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID := r.Context().Value(middleware.UserIDKey).(string)

	if id != userID {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = id

	if err := h.userService.Update(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Delete maneja la solicitud para eliminar un usuario.
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID := r.Context().Value(middleware.UserIDKey).(string)

	if id != userID {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}
	if err := h.userService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}