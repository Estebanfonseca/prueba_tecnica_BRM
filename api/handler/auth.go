package handler

import (
	"encoding/json"
	"net/http"
	"api_users/api/service"
	"api_users/api/models"
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


// Login godoc
// @Summary      Login
// @Description  Login a la aplicación y devuelve un token de acceso.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.LoginRequest true "Login request"
// @Success      200 {object} models.LoginResponse "Login successful"
// @Failure      400 {object} string "Invalid request body"
// @Failure      401 {object} string "Invalid credentials"
// @Router       /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Login maneja la solicitud de inicio de sesion y devuelve un token de acceso.
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	res := models.LoginResponse{
		Token: "Bearer " +token,
		}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
