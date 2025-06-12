package repository

import (
	"context"
	"api_users/api/models"
)

// UserRepository define la interfaz para las operaciones de usuarios.
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	List(ctx context.Context) ([]*models.UserResponse, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}