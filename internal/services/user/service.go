package user

import (
	"context"
	"github.com/vladislavninetyeight/service/internal/model"
	"github.com/vladislavninetyeight/service/internal/repositories"
	"github.com/vladislavninetyeight/service/internal/services"
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
	u, err := s.repo.Create(ctx, user)

	if err != nil {
		// TODO
	}

	return u.ID, nil
}
func (s *service) GetAll(ctx context.Context, filter *model.Filter) ([]model.User, error) {
	users, err := s.repo.GetAll(ctx, filter)

	if err != nil {
		// TODO
	}

	return users, nil
}
func (s *service) Get(ctx context.Context, id uint) (model.User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		// TODO
	}

	return user, nil
}
func (s *service) Update(ctx context.Context, detail model.UserDetail, id uint) (model.User, error) {
	err := s.repo.Update(ctx, id, detail)
	if err != nil {
		// TODO
	}

	user, err := s.Get(ctx, id)
	if err != nil {
		// TODO
	}

	return user, nil
}
func (s *service) Delete(ctx context.Context, id uint) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		// TODO
	}

	return nil
}
