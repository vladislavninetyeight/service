package converter

import (
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/model"
	rep "github.com/vladislavninetyeight/service/tree/main/internal/model/internal/repositories/user/model"
	"time"
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
		CreatedAT: time.Now(),
		UpdatedAT: time.Now(),
	}
}
