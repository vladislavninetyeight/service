package app

import (
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/providers"
)

type App struct {
	serviceProvider *providers.ServiceProvider
}

func New() (*App, error) {
	a := App{}
	err := a.init()

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *App) init() error {
	err := a.initServiceProvider()

	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider() error {
	a.serviceProvider = providers.NewServiceProvider()
	return nil
}

func (a *App) Serve() {
	a.serviceProvider.GetHTTPServer().Run()
}
