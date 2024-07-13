package user

import (
	"context"
	"errors"
	"github.com/vladislavninetyeight/service/internal/model"
	"github.com/vladislavninetyeight/service/internal/repositories"
	"github.com/vladislavninetyeight/service/internal/repositories/user/converter"
	rep "github.com/vladislavninetyeight/service/internal/repositories/user/model"
	"sort"
	"strings"
	"sync"
	"time"
)

var _ repositories.UserRepository = (*repository)(nil)

type repository struct {
	driver map[uint]*rep.User
	mu     sync.RWMutex
}

func NewUserRepository() *repository {
	return &repository{
		driver: make(map[uint]*rep.User),
		mu:     sync.RWMutex{},
	}
}

func (r *repository) GetAll(ctx context.Context, filter *model.Filter) ([]model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]rep.User, 0)

	r.applyFilters(&users, filter)
	r.applySorts(&users, filter)
	r.offset(&users, filter)
	r.limit(&users, filter)

	return converter.ToUsersFromReps(users), nil
}

func (r *repository) applyFilters(users *[]rep.User, filter *model.Filter) {
	for _, v := range r.driver {
		if filter.Name != "" && strings.ToLower(v.Detail.Name) != strings.ToLower(filter.Name) {
			continue
		}

		if filter.FromCreatedAt != nil && v.CreatedAT.Unix() < filter.FromCreatedAt.Unix() {
			continue
		}

		if filter.ToCreatedAt != nil && v.CreatedAT.Unix() > filter.ToCreatedAt.Unix() {
			continue
		}
		*users = append(*users, *v)
	}
}

func (r *repository) applySorts(users *[]rep.User, filter *model.Filter) {
	u := *users
	if strings.ToLower(filter.TopPostsAmount) == "asc" {
		sort.Slice(u, func(i, j int) bool {
			return u[i].ID < u[j].ID
		})
	} else {
		sort.Slice(u, func(i, j int) bool {
			return u[i].ID > u[j].ID
		})
	}
	*users = u
}

func (r *repository) offset(users *[]rep.User, filter *model.Filter) {
	if len(*users) < int(filter.Offset) {
		*users = make([]rep.User, 0)
	} else {
		*users = (*users)[filter.Offset:]
	}
}

func (r *repository) limit(users *[]rep.User, filter *model.Filter) {
	if filter.Limit != 0 && len(*users) > int(filter.Limit) {
		*users = (*users)[:filter.Limit]
	}
}

func (r *repository) Get(ctx context.Context, id uint) (model.User, error) {
	if _, ok := r.driver[id]; !ok {
		return model.User{}, errors.New("user not found")
	}

	return converter.ToUserFromRep(*r.driver[id]), nil
}

func (r *repository) Create(ctx context.Context, detail model.UserDetail) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user := converter.FromUserDetailToRep(uint(len(r.driver)), detail)
	user.CreatedAT = time.Now()
	user.UpdatedAT = time.Now()

	r.driver[user.ID] = &user

	return converter.ToUserFromRep(user), nil
}

func (r *repository) Update(ctx context.Context, id uint, user model.UserDetail) error {
	_, ok := r.driver[id]

	if !ok {
		return errors.New("user not found")
	}

	userRepDetail := converter.FromUserDetailToRepUserDetail(user)
	r.driver[id].Detail = userRepDetail
	r.driver[id].UpdatedAT = time.Now()

	return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	if _, ok := r.driver[id]; !ok {
		return errors.New("user not found")
	}

	delete(r.driver, id)

	return nil
}
