package repository

import (
	"context"
	"errors"

	"CRUD_API/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	CreateUser(ctx context.Context, req models.CreateUser) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}

type postgresUserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &postgresUserRepo{db: db}
}

func (r *postgresUserRepo) CreateUser(ctx context.Context, req models.CreateUser) (models.User, error) {
	query := `
		INSERT INTO users (name, email) 
		VALUES ($1, $2) 
		RETURNING id, name, email
	`

	var user models.User
	err := r.db.QueryRow(ctx, query, req.Name, req.Email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *postgresUserRepo) GetUsers(ctx context.Context) ([]models.User, error) {
	query := "SELECT id, name, email FROM users ORDER BY id DESC"

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
