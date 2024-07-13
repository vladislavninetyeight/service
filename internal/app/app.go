package app

import (
	"github.com/vladislavninetyeight/service/internal/app/router"
	"github.com/vladislavninetyeight/service/internal/providers"
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

func (a *App) Serve() error {
	routs := router.New(*a.serviceProvider)
	a.serviceProvider.GetHTTPServer().Handler = routs
	err := a.serviceProvider.GetHTTPServer().ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
