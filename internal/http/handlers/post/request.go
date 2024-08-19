package post

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/vladislavninetyeight/service/pkg/validator"
	"net/http"
	"net/url"
)

func StorePostValidate(request *http.Request) url.Values {
	rules := govalidator.MapData{
		"title":   []string{"min:5", "max:50"},
		"body":    []string{"min:20", "max:3000"},
		"user_id": []string{"numeric_between:1,"},
	}

	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request: request,
		Rules:   rules,
		Data:    &data,
	}

	return validator.New(opts).ValidateJSON()
}
