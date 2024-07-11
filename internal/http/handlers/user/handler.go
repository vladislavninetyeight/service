package user

import (
	"context"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/converter"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/model"
	"net/http"
)

type UserServiceProvider interface {
	Create(ctx context.Context, user model.UserDetail) (uint, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user model.UserDetail, id uint) (model.User, error)
}

func Store(request *http.Request, service UserServiceProvider) (uint, error) {
	ctx := context.TODO()
	
	detail := converter.FromRequestToUserDetail(request)
	userId, err := service.Create(ctx, detail)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func GetAll() {

}

func Update(id uint, detail model.UserDetail) {

}

func Delete(id uint) {

}
