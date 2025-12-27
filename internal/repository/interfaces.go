package repository

import (
	"context"

	"github.com/ghduuep/litly/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, b *domain.Book) (*domain.Book, error)
	FindByID(ctx context.Context, id int64) (*domain.Book, error)
	Delete(ctx context.Context, id int64) error
	FindAll(ctx context.Context) ([]*domain.Book, error)
	Update(ctx context.Context, b *domain.Book) error
}

type BookRepository interface {
	Create(ctx context.Context, b *domain.Book) (*domain.Book, error)
	FindByID(ctx context.Context, id int64) (*domain.Book, error)
	Delete(ctx context.Context, id int64) error
	FindAll(ctx context.Context) ([]*domain.Book, error)
	Update(ctx context.Context, b *domain.Book) error
}

type AuthorRepository interface {
	Create(ctx context.Context, b *domain.Book) (*domain.Author, error)
	FindByID(ctx context.Context, id int64) (*domain.Book, error)
	Delete(ctx context.Context, id int64) error
	FindAll(ctx context.Context) ([]*domain.Book, error)
	Update(ctx context.Context, b *domain.Book) error
}
