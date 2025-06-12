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

// Create godoc
// @Summary      Create User
// @Description  Create a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.UserRequest true "User object"
// @Success      201 {object} models.UserResponse "User created successfully"
// @Failure      400 {object} string "Invalid request body"
// @Failure      500 {object} string "Internal server error"
// @security      BearerAuth
// @Router       /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Create maneja la solicitud para crear un nuevo usuario.
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.userService.Create(r.Context(), &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := models.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}


// List godoc
// @Summary      List Users
// @Description  Get a list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200 {array} models.UserResponse "List of users"
// @Failure      500 {object} string "Internal server error"
// @security      BearerAuth
// @Router       /users [get]
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	// List maneja la solicitud para obtener una lista de todos los usuarios.
	users, err := h.userService.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}


// GetByID godoc
// @Summary      Get User by ID
// @Description  Get a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      200 {object} models.UserResponse "User found"
// @Failure      404 {object} string "User not found"
// @Failure      500 {object} string "Internal server error"
// @security     BearerAuth
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// GetByID maneja la solicitud para obtener un usuario por su ID.
	id := mux.Vars(r)["id"]

	user, err := h.userService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := models.UserResponse{
		ID:       user.ID,
		Name: user.Name,
		Email: user.Email,
	}
	json.NewEncoder(w).Encode(res)
}


// Update godoc
// @Summary      Update User
// @Description  Update an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Param        user body models.UserRequest true "User object"
// @Success      200 {object} models.UserResponse "User updated successfully"
// @Failure      400 {object} string "Invalid request body"
// @Failure      401 {object} string "Unauthorized access"
// @Failure      404 {object} string "User not found"
// @Failure      500 {object} string "Internal server error"
// @security      BearerAuth
// @Router       /users/{id} [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Update maneja la solicitud para actualizar un usuario existente.
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

	res := models.UserResponse{
		ID:       user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	json.NewEncoder(w).Encode(res)
}


// Delete godoc
// @Summary      Delete User
// @Description  Delete a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path string true "User ID"
// @Success      204 "User deleted successfully"
// @Failure      401 {object} string "Unauthorized access"
// @Failure      404 {object} string "User not found"
// @Failure      500 {object} string "Internal server error"
// @security      BearerAuth
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Delete maneja la solicitud para eliminar un usuario.
	id := mux.Vars(r)["id"]
	if err := h.userService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}