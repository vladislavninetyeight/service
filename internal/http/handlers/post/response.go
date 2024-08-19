package post

import "github.com/vladislavninetyeight/service/internal/model"

type GetAllPostResponse struct {
	Posts []model.Post
}

func FromPostsToGetAllResponse(posts []model.Post) GetAllPostResponse {
	return GetAllPostResponse{
		Posts: posts,
	}
}
