package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ghduuep/litly/internal/domain"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email, created_at, updated_at`

	newUser := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, u.Username, u.Email, u.Password).Scan(
		&newUser.ID,
		&newUser.Username,
		&newUser.Email,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1`

	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found with id %d", id)
		}
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, id int64, u *domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *PostgresUserRepository) findAll(ctx context.Context) ([]*domain.User, error) {
	return nil, nil
}
