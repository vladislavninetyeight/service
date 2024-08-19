package post

import (
	"context"
	"github.com/vladislavninetyeight/service/internal/converter"
	"github.com/vladislavninetyeight/service/internal/model"
	"net/http"
)

type Service interface {
	Create(ctx context.Context, post model.PostDetail) error
	GetAll(ctx context.Context) ([]model.Post, error)
}

func Store(request *http.Request, service Service) error {
	ctx := context.TODO()

	post := converter.FromRequestToPostDetail(request)
	err := service.Create(ctx, post)
	if err != nil {
		// TODO
	}

	return nil
}

func GetAll(request *http.Request, service Service) (GetAllPostResponse, error) {
	ctx := context.TODO()

	posts, err := service.GetAll(ctx)
	if err != nil {
		//TODO
	}

	return FromPostsToGetAllResponse(posts), nil
}
