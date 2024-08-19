package user

import "github.com/vladislavninetyeight/service/internal/model"

type GetAllUserResponse struct {
	Users []model.User
}

type UpdateUserResponse struct {
	model.User
}

func FromUsersToGetAllUserResponse(users []model.User) GetAllUserResponse {
	return GetAllUserResponse{
		Users: users,
	}
}

func FromUserToUpdateUserResponse(user model.User) UpdateUserResponse {
	return UpdateUserResponse{
		User: user,
	}
}
