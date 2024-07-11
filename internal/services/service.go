package services

import (
	"context"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/model"
)

type UserService interface {
	Create(ctx context.Context, user model.UserDetail) (uint, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user model.UserDetail, id uint) (model.User, error)
}

type PostService interface {
	Create(ctx context.Context, post model.PostDetail) (uint, error)
	GetAll(ctx context.Context) ([]model.Post, error)
	Update(ctx context.Context, post model.PostDetail, id uint) (model.Post, error)
}
