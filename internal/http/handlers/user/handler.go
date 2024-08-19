package user

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/vladislavninetyeight/service/internal/converter"
	"github.com/vladislavninetyeight/service/internal/model"
	"net/http"
	"strconv"
)

type Service interface {
	Create(ctx context.Context, user model.UserDetail) (uint, error)
	GetAll(ctx context.Context, filter *model.Filter) ([]model.User, error)
	Update(ctx context.Context, user model.UserDetail, id uint) (model.User, error)
	Delete(ctx context.Context, id uint) error
}

func Store(request *http.Request, service Service) (uint, error) {
	ctx := context.TODO()

	detail := converter.FromRequestToUserDetail(request)
	userId, err := service.Create(ctx, detail)

	if err != nil {
		// TODO
	}

	return userId, nil
}

func GetAll(request *http.Request, service Service) (GetAllUserResponse, error) {
	ctx := context.TODO()

	filters, err := converter.FromRequestToFilter(request)
	if err != nil {
		// TODO
	}

	users, err := service.GetAll(ctx, filters)
	if err != nil {
		// TODO
	}

	return FromUsersToGetAllUserResponse(users), nil
}

func Update(request *http.Request, service Service) (UpdateUserResponse, error) {
	id := chi.URLParam(request, "id")
	ctx := context.TODO()

	userId, err := strconv.Atoi(id)
	if err != nil {
		// TODO
	}
	detail := converter.FromRequestToUserDetail(request)

	user, err := service.Update(ctx, detail, uint(userId))
	if err != nil {
		// TODO
	}

	return FromUserToUpdateUserResponse(user), nil
}

func Delete(request *http.Request, service Service) error {
	id := chi.URLParam(request, "id")
	ctx := context.TODO()

	userId, err := strconv.Atoi(id)
	if err != nil {
		// TODO
	}

	err = service.Delete(ctx, uint(userId))
	if err != nil {
		// TODO
	}

	return nil
}
