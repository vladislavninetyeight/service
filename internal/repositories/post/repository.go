package post

import (
	"context"
	"github.com/vladislavninetyeight/service/internal/model"
	"github.com/vladislavninetyeight/service/internal/repositories"
	rep "github.com/vladislavninetyeight/service/internal/repositories/post/model"
	"sync"
)

var _ repositories.PostRepository = (*repository)(nil)

type repository struct {
	driver map[uint]*rep.Post
	mu     sync.RWMutex
}

func NewPostRepository() *repository {
	return &repository{
		driver: make(map[uint]*rep.Post),
		mu:     sync.RWMutex{},
	}
}

func (r *repository) Create(ctx context.Context, post model.PostDetail) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetAll(ctx context.Context) ([]model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Get(ctx context.Context, id uint) (model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, post model.PostDetail, id uint) (model.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
