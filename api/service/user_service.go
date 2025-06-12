package service

import (
	"api_users/api/models"
	"api_users/api/repository"
	"context"

	"golang.org/x/crypto/bcrypt"
)

// UserService define el modelo de servicio de usuario.
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService crea una nueva instancia de UserService.
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Create crea un nuevo usuario lo valida y lo guarda en la base de datos.
func (s *UserService) Create(ctx context.Context, user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Create(ctx, user)
}

// List devuelve una lista de todos los usuarios.
func (s *UserService) List(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.List(ctx)
}

// GetByID obtiene un usuario por su ID.
func (s *UserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

// Update actualiza un usuario existente.
func (s *UserService) Update(ctx context.Context, user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.userRepo.Update(ctx, user)

}

// Delete elimina un usuario.
func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.userRepo.Delete(ctx, id)
}