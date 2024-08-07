package services

import (
	"context"
	"github.com/vladislavninetyeight/service/internal/model"
)

type UserService interface {
	Create(ctx context.Context, user model.UserDetail) (uint, error)
	GetAll(ctx context.Context, filter *model.Filter) ([]model.User, error)
	Update(ctx context.Context, user model.UserDetail, id uint) (model.User, error)
	Delete(ctx context.Context, id uint) error
}

type PostService interface {
	Create(ctx context.Context, post model.PostDetail) (uint, error)
	GetAll(ctx context.Context) ([]model.Post, error)
	Update(ctx context.Context, post model.PostDetail, id uint) (model.Post, error)
	Delete(ctx context.Context, id uint) error
}
