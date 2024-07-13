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
		panic(err)
	}
	return detail
}

func FromUsersToGetAllUserResponse(users []model.User) model.GetAllUserResponse {
	return model.GetAllUserResponse{
		Users: users,
	}
}

func FromUserToUpdateUserResponse(user model.User) model.UpdateUserResponse {
	return model.UpdateUserResponse{
		User: user,
	}
}
