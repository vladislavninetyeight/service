package post

import (
	"context"
	"github.com/vladislavninetyeight/service/internal/model"
	"github.com/vladislavninetyeight/service/internal/repositories"
	"github.com/vladislavninetyeight/service/internal/repositories/post/converter"
	rep "github.com/vladislavninetyeight/service/internal/repositories/post/model"
	"sync"
	"time"
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

func (r *repository) Create(ctx context.Context, detail model.PostDetail) (uint, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	post := converter.FromPostDetailToRep(uint(len(r.driver)), detail)
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	r.driver[post.ID] = &post

	return post.ID, nil
}

func (r *repository) GetAll(ctx context.Context) ([]model.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	posts := make([]rep.Post, 0)

	for _, v := range r.driver {
		posts = append(posts, *v)
	}

	return converter.ToPostsFromReps(posts), nil
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
