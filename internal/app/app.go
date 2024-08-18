package app

import (
	"github.com/vladislavninetyeight/service/internal/app/router"
	"github.com/vladislavninetyeight/service/internal/app/server"
	"github.com/vladislavninetyeight/service/internal/config"
	"net/http"
)

type App struct {
	Server *http.Server
}

func New(config *config.Config) *App {
	app := &App{}
	app.initServer(&config.Server)
	return app
}

func (a *App) initServer(config *config.Server) {
	routs := router.New()
	a.Server = server.New(config)
	a.Server.Handler = routs
}

func (a *App) Serve() error {
	err := a.Server.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
