package post

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/vladislavninetyeight/service/internal/model"
	"github.com/vladislavninetyeight/service/pkg/validator"
	"net/url"
)

func StorePostValidate(post model.PostDetail) url.Values {
	rules := govalidator.MapData{
		"title":   []string{"min:5", "max:50"},
		"body":    []string{"min:20", "max:3000"},
		"user_id": []string{"numeric_between:1,"},
	}

	opts := govalidator.Options{
		Rules: rules,
		Data:  &post,
	}

	return validator.New(opts).ValidateStruct()
}
