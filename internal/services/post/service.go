package user

import (
	"context"
	"github.com/vladislavninetyeight/service/internal/model"
	"github.com/vladislavninetyeight/service/internal/repositories"
	"github.com/vladislavninetyeight/service/internal/services"
)

var _ services.PostService = (*service)(nil)

type service struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, post model.PostDetail) (uint, error) {
	id, err := s.repo.Create(ctx, post)

	if err != nil {
		// TODO
	}

	return id, nil
}
func (s *service) GetAll(ctx context.Context) ([]model.Post, error) {
	return []model.Post{}, nil
}
func (s *service) Get(ctx context.Context, id uint) (model.Post, error) {
	return model.Post{}, nil
}
func (s *service) Update(ctx context.Context, post model.PostDetail, id uint) (model.Post, error) {
	return model.Post{}, nil
}
func (s *service) Delete(ctx context.Context, id uint) error {
	return nil
}
