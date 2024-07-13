package router

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vladislavninetyeight/service/internal/http/handlers/user"
	"github.com/vladislavninetyeight/service/internal/providers"
	"net/http"
)

func New(provider providers.ServiceProvider) chi.Router {
	router := chi.NewRouter()
	return initRoutes(router, provider)
}

func initRoutes(r chi.Router, provider providers.ServiceProvider) chi.Router {
	r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
		id, err := user.Store(request, provider.GetUserService())
		if err != nil {
			// TODO
		}

		response, err := json.Marshal(id)
		if err != nil {
			// TODO
		}

		_, err = writer.Write(response)
		if err != nil {
			// TODO
		}
	})

	r.Get("/users", func(writer http.ResponseWriter, request *http.Request) {
		users, err := user.GetAll(request, provider.GetUserService())
		if err != nil {
			// TODO
		}

		response, err := json.Marshal(users)
		if err != nil {
			// TODO
		}

		_, err = writer.Write(response)
		if err != nil {
			// TODO
		}
	})

	r.Patch("/user/{id}", func(writer http.ResponseWriter, request *http.Request) {
		u, err := user.Update(request, provider.GetUserService())
		if err != nil {
			// TODO
		}

		response, err := json.Marshal(u)
		if err != nil {
			// TODO
		}

		_, err = writer.Write(response)
		if err != nil {
			// TODO
		}
	})

	return r
}
