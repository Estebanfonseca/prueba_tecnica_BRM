package handler

import (
	"encoding/json"
	"net/http"
	"api_users/api/service"
)

// AuthHandler maneja las solicitudes de autenticación.
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler crea una nueva instancia de AuthHandler.
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// LoginRequest define la estructura de la solicitud de inicio de sesion.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse define la estructura de la respuesta de inicio de sesion.
type LoginResponse struct {
	Token string `json:"token"`
}

// Login maneja la solicitud de inicio de sesion y devuelve un token de acceso.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	res := LoginResponse{
		Token: token,
		}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
