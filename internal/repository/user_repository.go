package repository

import (
	"context"
	"mathalama/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, username string, email string) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
}

type postgresUserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &postgresUserRepository{pool: pool}
}

func (r *postgresUserRepository) CreateUser(ctx context.Context, username string, email string) error {
	query := `INSERT INTO users (username, email) VALUES ($1, $2)`
	_, err := r.pool.Exec(ctx, query, username, email)
	return err
}

func (r *postgresUserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := `SELECT id, username, email FROM users WHERE username = $1`
	row := r.pool.QueryRow(ctx, query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *postgresUserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, username, email FROM users`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}
