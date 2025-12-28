package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ghduuep/litly/internal/domain"
	dto "github.com/ghduuep/litly/internal/dto/user"
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

func (r *PostgresUserRepository) Update(ctx context.Context, id int64, u *dto.UpdateUserRequest) (*domain.User, error) {
	var updates []string
	var args []any
	argCounter := 1

	if u.Username != nil {
		updates = append(updates, fmt.Sprintf("name = $%d", argCounter))
		args = append(args, *u.Username)
		argCounter++
	}

	if u.Email != nil {
		updates = append(updates, fmt.Sprintf("email = $%d", argCounter))
		args = append(args, *u.Email)
		argCounter++
	}

	if u.Password != nil {
		updates = append(updates, fmt.Sprintf("password = $%d", argCounter))
		args = append(args, *u.Password)
		argCounter++
	}

	args = append(args, id)

	query := fmt.Sprintf(`
			UPDATE users
			SET %s
			WHERE id = $%d
			RETURNING id, username, email, created_at, updated_at`,
		strings.Join(updates, ", "),
		argCounter,
	)

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

func (r *PostgresUserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *PostgresUserRepository) findAll(ctx context.Context) ([]*domain.User, error) {
	query := `SELECT id, username, email, created_at, updated_at FROM users`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
