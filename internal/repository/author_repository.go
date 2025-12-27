package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ghduuep/litly/internal/domain"
)

type PostgresAuthorRepository struct {
	db *sql.DB
}

func NewPostgresAuthorRepository(db *sql.DB) *PostgresAuthorRepository {
	return &PostgresAuthorRepository{
		db: db,
	}
}

func (r *PostgresAuthorRepository) Create(ctx context.Context, a *domain.Author) (*domain.Author, error) {
	query := `INSERT INTO users (name, description) VALUES ($1, $2) RETURNING id, name, description, created_at, updated_at`

	newAuthor := &domain.Author{}
	err := r.db.QueryRowContext(ctx, query, a.Name, a.Description).Scan(
		&newAuthor.ID,
		&newAuthor.Name,
		&newAuthor.Description,
		&newAuthor.CreatedAt,
		&newAuthor.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return newAuthor, nil
}

func (r *PostgresAuthorRepository) FindByID(ctx context.Context, id int64) (*domain.Author, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM users WHERE id = $1`

	author := &domain.Author{}
	err := r.db.QueryRowContext(ctx, query).Scan(
		&author.ID,
		&author.Name,
		&author.Description,
		&author.CreatedAt,
		&author.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found with id %d", id)
		}
		return nil, err
	}

	return author, nil
}

func (r *PostgresAuthorRepository) Update(ctx context.Context, id int64, u *domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *PostgresAuthorRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM authors WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *PostgresAuthorRepository) findAll(ctx context.Context) ([]*domain.User, error) {
	return nil, nil
}
