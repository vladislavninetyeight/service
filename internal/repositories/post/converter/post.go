package converter

import (
	"github.com/vladislavninetyeight/service/internal/model"
	rep "github.com/vladislavninetyeight/service/internal/repositories/post/model"
)

func FromPostDetailToRep(id uint, detail model.PostDetail) rep.Post {
	return rep.Post{
		ID: id,
		Detail: rep.PostDetail{
			Title:  detail.Title,
			Body:   detail.Body,
			UserID: detail.UserID,
		},
	}
}
