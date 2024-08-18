package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/vladislavninetyeight/service/internal/model"
	"sync"
	"time"
)

type Repository interface {
	Create(ctx context.Context, user model.UserDetail) (model.User, error)
	GetAll(ctx context.Context, filter *model.Filter) ([]model.User, error)
	Get(ctx context.Context, id uint) (model.User, error)
	Update(ctx context.Context, id uint, user model.UserDetail) error
	Delete(ctx context.Context, id uint) error
}

type Driver interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, arguments ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, arguments ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

var _ Repository = (*repository)(nil)

type repository struct {
	driver Driver
	mu     sync.RWMutex
}

func NewUserRepository(driver Driver) *repository {
	return &repository{
		driver: driver,
		mu:     sync.RWMutex{},
	}
}

func (r *repository) GetAll(ctx context.Context, filter *model.Filter) ([]model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	query := `
		SELECT *
		FROM users
	`
	params := make([]any, 0, 6)

	r.applyFilters(&query, &params, filter)
	r.applySorts(&query, &params, filter)
	r.offset(&query, &params, filter)
	r.limit(&query, &params, filter)

	users, err := r.driver.Query(ctx, query, params...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)

	return []model.User{}, nil
}

func (r *repository) applyFilters(query *string, params *[]any, filter *model.Filter) {
	*query += ` WHERE (LOWER(name) = LOWER($1) OR $1 = '')
				AND (created_at < $2 OR $2 = '')
				AND (created_at > $3 OR $3 = '')
	`
	*params = append(*params, filter.Name, filter.FromCreatedAt, filter.ToCreatedAt)
}

func (r *repository) applySorts(query *string, params *[]any, filter *model.Filter) {
	*query += ` ORDER BY created_at $4
	`
	*params = append(*params, filter.Sort)
}

func (r *repository) offset(query *string, params *[]any, filter *model.Filter) {
	*query += ` OFFSET $5
	`
	*params = append(*params, filter.Offset)
}

func (r *repository) limit(query *string, params *[]any, filter *model.Filter) {
	*query += ` LIMIT $6
	`
	*params = append(*params, filter.Limit)
}

func (r *repository) Get(ctx context.Context, id uint) (user model.User, err error) {
	query := "SELECT * FROM users WHERE id=?"
	pgx, err := r.driver.Query(ctx, query, id)
	if err != nil {
		return model.User{}, err
	}
	val, err := pgx.Values()
	if err != nil {
		return model.User{}, err
	}
	fmt.Println(val)

	return model.User{}, nil
}

func (r *repository) Create(ctx context.Context, detail model.UserDetail) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	sql := `INSERT INTO "user" (name, phone, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	_, err := r.driver.Exec(ctx, sql, detail.Name, detail.Phone,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02 15:04:05"))

	if err != nil {
		return model.User{}, err
	}

	return model.User{}, nil
}

func (r *repository) Update(ctx context.Context, id uint, user model.UserDetail) error {
	return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	return nil
}
