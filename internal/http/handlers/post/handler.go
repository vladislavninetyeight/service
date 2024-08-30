package post

import (
	"context"
	"fmt"
	"github.com/vladislavninetyeight/service/internal/model"
	"net/http"
)

type Service interface {
	Create(ctx context.Context, post model.PostDetail) error
	GetAll(ctx context.Context) ([]model.Post, error)
}

func Store(_ *http.Request, detail model.PostDetail, service Service) error {
	ctx := context.TODO()

	err := service.Create(ctx, detail)
	fmt.Println(err)
	if err != nil {
		// TODO
	}

	return nil
}

func GetAll(_ *http.Request, service Service) (GetAllPostResponse, error) {
	ctx := context.TODO()

	posts, err := service.GetAll(ctx)
	if err != nil {
		//TODO
	}

	return FromPostsToGetAllResponse(posts), nil
}
