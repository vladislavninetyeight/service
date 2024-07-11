package router

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/http/handlers/user"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/providers"
	"log"
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
			panic(err)
		}
		response, err := json.Marshal(id)
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(response)
		if err != nil {
			log.Fatal(err)
		}
	})

	return r
}
