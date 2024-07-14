package converter

import (
	"encoding/json"
	"github.com/vladislavninetyeight/service/internal/model"
	"net/http"
)

func FromRequestToPostDetail(request *http.Request) model.PostDetail {
	var detail model.PostDetail

	err := json.NewDecoder(request.Body).Decode(&detail)
	if err != nil {
		// TODO
	}
	return detail
}

func FromPostsToGetAllResponse(posts []model.Post) model.PostGetAllResponse {
	return model.PostGetAllResponse{
		Posts: posts,
	}
}
