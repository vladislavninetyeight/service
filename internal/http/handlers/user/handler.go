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

func GetAll(request *http.Request, service Service) (model.GetAllUserResponse, error) {
	ctx := context.TODO()

	filters, err := converter.FromRequestToFilter(request)
	if err != nil {
		// TODO
	}

	users, err := service.GetAll(ctx, filters)
	if err != nil {
		// TODO
	}

	return converter.FromUsersToGetAllUserResponse(users), nil
}

func Update(request *http.Request, service Service) (model.UpdateUserResponse, error) {
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

	return converter.FromUserToUpdateUserResponse(user), nil
}

func Delete(id uint) {

}
