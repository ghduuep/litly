package service

import (
	"context"
	"time"

	"github.com/ghduuep/litly/internal/domain"
	"github.com/ghduuep/litly/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *UserService) FindAll(ctx context.Context) ([]*domain.User, error) {

}
