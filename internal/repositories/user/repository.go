package user

import (
	"context"
	"errors"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/model"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/repositories"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/repositories/user/converter"
	rep "github.com/vladislavninetyeight/service/tree/main/internal/model/internal/repositories/user/model"
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

func (r *repository) GetAll(ctx context.Context) ([]model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]rep.User, 0, len(r.driver))

	for _, v := range r.driver {
		users = append(users, *v)
	}

	return converter.ToUsersFromReps(users), nil
}

func (r *repository) Get(ctx context.Context, id uint) (model.User, error) {
	if _, ok := r.driver[id]; ok {
		return model.User{}, errors.New("user not found")
	}

	return converter.ToUserFromRep(*r.driver[id]), nil
}

func (r *repository) Create(ctx context.Context, detail model.UserDetail) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user := converter.FromUserDetailToRep(uint(len(r.driver)), detail)
	r.driver[user.ID] = &user

	return converter.ToUserFromRep(user), nil
}

func (r *repository) Update(ctx context.Context, id uint, user model.UserDetail) error {
	_, ok := r.driver[id]

	if !ok {
		return errors.New("user not found")
	}

	userRep := converter.FromUserDetailToRep(id, user)
	userRep.UpdatedAT = time.Now()
	*r.driver[id] = userRep
	return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	if _, ok := r.driver[id]; !ok {
		return errors.New("user not found")
	}

	delete(r.driver, id)

	return nil
}
