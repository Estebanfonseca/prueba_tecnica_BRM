package models

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// User representa un usuario.
type User struct {
	ID string `json:"id" example:"12345"`
	Name string `json:"name" example:"user"`
	Email string `json:"email" example:"email@example.com"`
	Password string `json:"password" example:"password123"`
}

// UserRequest representa el objeto de peticion de usuario.
type UserRequest struct {
	Name    string `json:"name" example:"example"`
	Email   string `json:"email" example:"email@example.com"`
	Password string `json:"password" example:"password123"`
}

// UserResponse representa el objeto de respuesta de usuario.
type UserResponse struct {
	ID string   `json:"id" example:"1abc"`
	Name string `json:"name" example:"example"`
	Email string `json:"email" example:"email@example.com"`
}

// LoginRequest representa el objeto de peticion de inicio de sesion
type LoginRequest struct {
	Email    string `json:"email" example:"email@example.com"`
	Password string `json:"password" example:"password123"`
}

// Login Response representa el objeto de respuesta de inicio de sesion
type LoginResponse struct {
	Token string `json:"token" example:"token123"`
}


// Validate verifica que los campos requeridos del usuario no estén vacíos.
// Si algún campo requerido está vacío, devuelve un error.
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("nombre es requerido")
	}
	if u.Email == "" {
		return errors.New("email es requerido")
	}
	if u.Password == "" {
		return errors.New("constraseña es requerida")
	}
	return nil
}

// Claims representa los datos del usuario en el token JWT.
type Claims struct {
	UserID string `json:"user_id" example:"12345"`
	Email string `json:"email" example:"email@example.com"`
	jwt.RegisteredClaims
}

// UserClaims crea un nuevo objeto Claims con los datos del usuario y la duración del token.
func UserClaims(userID, email string, duration time.Duration) *Claims {
	return &Claims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "api_users",
		},
	}
}
