package main

import (
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		panic(err)
	}

	application.Serve()
}
