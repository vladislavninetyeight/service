package converter

import (
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/model"
	"net/http"
)

func FromRequestToUserDetail(request *http.Request) model.UserDetail {
	return model.UserDetail{
		Name:  request.FormValue("name"),
		Phone: request.FormValue("surname"),
	}
}
