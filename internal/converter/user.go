package converter

import (
	"encoding/json"
	"github.com/vladislavninetyeight/service/internal/model"
	"net/http"
)

func FromRequestToUserDetail(request *http.Request) model.UserDetail {
	var detail model.UserDetail

	err := json.NewDecoder(request.Body).Decode(&detail)
	if err != nil {
		// TODO
	}
	return detail
}
