package converter

import (
	"github.com/vladislavninetyeight/service/internal/model"
	rep "github.com/vladislavninetyeight/service/internal/repositories/post/model"
)

func ToPostFromRep(post rep.Post) model.Post {
	return model.Post{
		ID: post.ID,
		Detail: model.PostDetail{
			Title:  post.Detail.Title,
			Body:   post.Detail.Body,
			UserID: post.Detail.UserID,
		},
		CreatedAT: post.CreatedAt,
		UpdatedAT: post.UpdatedAt,
	}
}

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

func ToPostsFromReps(reps []rep.Post) []model.Post {
	posts := make([]model.Post, 0)
	for _, post := range reps {
		posts = append(posts, ToPostFromRep(post))
	}

	return posts
}
