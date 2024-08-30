package post

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/vladislavninetyeight/service/internal/model"
	"github.com/vladislavninetyeight/service/internal/repositories/post/converter"
	rep "github.com/vladislavninetyeight/service/internal/repositories/post/model"
	"sync"
	"time"
)

type Repository interface {
	Create(ctx context.Context, detail model.PostDetail) (post model.Post, err error)
	GetAll(ctx context.Context) ([]model.Post, error)
	Get(ctx context.Context, id uint) (model.Post, error)
	Update(ctx context.Context, post model.PostDetail, id uint) (model.Post, error)
	Delete(ctx context.Context, id uint) error
}

type Driver interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

var _ = (*repository)(nil)

type repository struct {
	driver Driver
	mu     sync.RWMutex
}

func NewPostRepository(driver Driver) *repository {
	return &repository{
		driver: driver,
		mu:     sync.RWMutex{},
	}
}

func (r *repository) Create(ctx context.Context, detail model.PostDetail) (post model.Post, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var repPost rep.Post

	sql := `INSERT INTO post (title, body, user_id, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING post_id, title, body, user_id, created_at, updated_at`

	row := r.driver.QueryRow(ctx, sql, detail.Title, detail.Body, detail.UserID, time.Now(), time.Now())

	err = row.Scan(&repPost.ID, &repPost.Detail.Title, &repPost.Detail.Body, &repPost.Detail.UserID, &repPost.CreatedAt, &repPost.UpdatedAt)
	if err != nil {
		return
	}

	post = converter.ToPostFromRep(repPost)
	return
}

func (r *repository) GetAll(ctx context.Context) ([]model.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	sql := "SELECT * FROM posts"

	_, err := r.driver.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	return converter.ToPostsFromReps(), nil
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
