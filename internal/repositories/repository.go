package repositories

import (
	"context"
	"github.com/vladislavninetyeight/service/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.UserDetail) (model.User, error)
	GetAll(ctx context.Context, filter *model.Filter) ([]model.User, error)
	Get(ctx context.Context, id uint) (model.User, error)
	Update(ctx context.Context, id uint, user model.UserDetail) error
	Delete(ctx context.Context, id uint) error
}

type PostRepository interface {
	Create(ctx context.Context, post model.PostDetail) (uint, error)
	GetAll(ctx context.Context) ([]model.Post, error)
	Get(ctx context.Context, id uint) (model.Post, error)
	Update(ctx context.Context, post model.PostDetail, id uint) (model.Post, error)
	Delete(ctx context.Context, id uint) error
}
