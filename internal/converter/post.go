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

func ToPostsFromReps() []model.Post {
	return []model.Post{}
}
