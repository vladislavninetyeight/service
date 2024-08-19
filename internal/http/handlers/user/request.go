package user

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/vladislavninetyeight/service/pkg/validator"
	"net/http"
	"net/url"
)

func StoreUserValidate(request *http.Request) url.Values {
	rules := govalidator.MapData{
		"name":  []string{"required", "min:2"},
		"phone": []string{"required", "min:11", "max:11"},
	}

	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request: request,
		Rules:   rules,
		Data:    &data,
	}

	return validator.New(opts).ValidateJSON()
}

func GetUsersValidate(request *http.Request) url.Values {

	rules := govalidator.MapData{
		"fromCreatedAt": []string{"timestamp"},
		"toCreatedAt":   []string{"timestamp"},
		"name":          []string{"min:2"},
		"limit":         []string{"numeric_between:1,"},
		"offset":        []string{"numeric_between:1,"},
		"sort":          []string{"order"},
	}

	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request: request,
		Rules:   rules,
		Data:    &data,
	}

	return validator.New(opts).Validate()
}

func UpdateUserValidate(request *http.Request) url.Values {
	rules := govalidator.MapData{
		"name":  []string{"min:2"},
		"phone": []string{"min:11", "max:11"},
	}

	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request: request,
		Rules:   rules,
		Data:    &data,
	}

	return validator.New(opts).ValidateJSON()
}
