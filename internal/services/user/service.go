package user

import (
	"context"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/model"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/repositories"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/services"
)

var _ services.UserService = (*service)(nil)

type service struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, user model.UserDetail) (uint, error) {
	return 0, nil
}
func (s *service) GetAll(ctx context.Context) ([]model.User, error) {
	return []model.User{}, nil
}
func (s *service) Get(ctx context.Context, id uint) (model.User, error) {
	return model.User{}, nil
}
func (s *service) Update(ctx context.Context, user model.UserDetail, id uint) (model.User, error) {
	return model.User{}, nil
}
func (s *service) Delete(ctx context.Context, id uint) error {
	return nil
}
