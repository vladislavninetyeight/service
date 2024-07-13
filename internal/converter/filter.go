package converter

import (
	"fmt"
	"github.com/vladislavninetyeight/service/internal/model"
	"net/http"
	"strconv"
	"time"
)

func FromRequestToFilter(request *http.Request) (*model.Filter, error) {
	fromCreatedAt, err := getFilterTime(request.URL.Query().Get("fromCreatedAt"))
	if err != nil {
		// TODO
	}

	toCreatedAt, err := getFilterTime(request.URL.Query().Get("toCreatedAt"))
	if err != nil {
		// TODO
	}

	limit, err := strconv.Atoi(request.URL.Query().Get("limit"))
	if err != nil {
		// TODO
	}

	offset, err := strconv.Atoi(request.URL.Query().Get("offset"))
	if err != nil {
		// TODO
	}

	return &model.Filter{
		FromCreatedAt:  fromCreatedAt,
		ToCreatedAt:    toCreatedAt,
		Name:           request.URL.Query().Get("name"),
		Limit:          uint(limit),
		Offset:         uint(offset),
		TopPostsAmount: request.URL.Query().Get("topPostsAmount"),
	}, nil
}

func getFilterTime(val string) (*time.Time, error) {
	var err error
	var timeAt time.Time

	if val != "" {
		timeAt, err = time.Parse(time.RFC3339, val)
		if err != nil {
			return nil, fmt.Errorf("invalid value for from_created_at: %v", err)
		}
		fmt.Println(timeAt)
		return &timeAt, nil
	}

	return nil, nil
}
