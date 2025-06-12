package repository

import (
	"api_users/api/models"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// PostgresRepository es la instancia de la conexion a base de datos.
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository crea una nueva instancia de PostgresRepository.
func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Create inserta un nuevo usuario en la base de datos.
func (r *PostgresRepository) Create(ctx context.Context, user *models.User) error {
	user.ID = uuid.New().String()
	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Password)
	return err
}

// List devuelve una lista de todos los usuarios en la base de datos.
func (r *PostgresRepository) List(ctx context.Context) ([]*models.UserResponse, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.UserResponse
	for rows.Next() {
		user := &models.UserResponse{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByID obtiene un usuario por su ID.
func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		return nil, errors.New("usuario no encontrado")
	}

	return user, nil
}

// Update actualiza un usuario existente en la base de datos.
func (r *PostgresRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.ID)
	return err
}

// Delete elimina un usuario de la base de datos.
func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GetByEmail obtiene un usuario por su correo electrónico.
func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return nil , errors.New("usuario no encontrado")

	}
	return user, err
}