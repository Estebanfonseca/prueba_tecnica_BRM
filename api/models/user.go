package models

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// User representa un usuario.
type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
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
	UserID string `json:"user_id"`
	Email string `json:"email"`
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
