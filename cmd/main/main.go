package main

import (
	"github.com/vladislavninetyeight/service/internal/app"
	"github.com/vladislavninetyeight/service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	application := app.New(cfg)

	err := application.Serve()
	if err != nil {
		panic(err)
	}
}
