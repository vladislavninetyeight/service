package main

import (
	"github.com/vladislavninetyeight/service/internal/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		panic(err)
	}

	err = application.Serve()
	if err != nil {
		panic(err)
	}
}
