package service

import (
	"context"
	"errors"
	"time"
	"api_users/api/models"
	"api_users/api/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Authservice define el modelo de servicio de autenticación.
type AuthService struct {
	userRepo repository.UserRepository
	secretKey string
	tokenDuration time.Duration
}

// NewAuthService crea una nueva instancia de AuthService.
func NewAuthService(userRepo repository.UserRepository, secretKey string, tokenDuration time.Duration) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		secretKey: secretKey,
		tokenDuration: tokenDuration,
	}
}

// Login inicia sesion y genera el token de acceso.
func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("credenciales inválidas")
	}

	claims := models.UserClaims(user.ID, user.Email, s.tokenDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenLogin, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", errors.New("error al generar el token")
	}

	return tokenLogin, nil
}

