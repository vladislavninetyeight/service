package converter

import (
	"github.com/vladislavninetyeight/service/internal/model"
	rep "github.com/vladislavninetyeight/service/internal/repositories/user/model"
)

func ToUserFromRep(user rep.User) model.User {
	return model.User{
		ID: user.ID,
		Detail: model.UserDetail{
			Name:  user.Detail.Name,
			Phone: user.Detail.Phone,
		},
		CreatedAT: user.CreatedAT,
		UpdatedAT: user.UpdatedAT,
	}
}

func ToUsersFromReps(reps []rep.User) []model.User {
	users := make([]model.User, 0)
	for _, user := range reps {
		users = append(users, ToUserFromRep(user))
	}

	return users
}

func FromUserDetailToRep(id uint, detail model.UserDetail) rep.User {
	return rep.User{
		ID: id,
		Detail: rep.UserDetail{
			Name:  detail.Name,
			Phone: detail.Phone,
		},
	}
}

func FromUserDetailToRepUserDetail(detail model.UserDetail) rep.UserDetail {
	return rep.UserDetail{
		Name:  detail.Name,
		Phone: detail.Phone,
	}
}
