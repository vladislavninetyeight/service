package router

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/vladislavninetyeight/service/internal/http/handlers/post"
	"github.com/vladislavninetyeight/service/internal/http/handlers/user"
	"github.com/vladislavninetyeight/service/internal/http/middleware"
	"github.com/vladislavninetyeight/service/internal/http/requests"
	"github.com/vladislavninetyeight/service/internal/providers"
	"net/http"
	"net/url"
)

func New() chi.Router {
	router := chi.NewRouter()
	return initRoutes(router)
}

func initRoutes(r chi.Router) chi.Router {
	provider := providers.Container()
	r.Use(middleware.RequestLogger())

	r.Group(func(r chi.Router) {
		r.Post("/user", func(writer http.ResponseWriter, request *http.Request) {
			errMessages := requests.StoreUserValidate(request)
			if errMessages.Encode() != "" {
				err := badRequest(writer, errMessages)
				if err != nil {
					panic(err)
				}
				return
			}

			id, err := user.Store(request, provider.GetUserService())
			if err != nil {
				// TODO
			}

			response, err := json.Marshal(id)
			if err != nil {
				// TODO
			}

			writer.WriteHeader(http.StatusCreated)
			_, err = writer.Write(response)
			if err != nil {
				// TODO
			}
		})

		r.Get("/users", func(writer http.ResponseWriter, request *http.Request) {
			errMessages := requests.GetUsersValidate(request)

			if errMessages.Encode() != "" {
				err := badRequest(writer, errMessages)
				if err != nil {
					panic(err)
				}
				return
			}

			users, err := user.GetAll(request, provider.GetUserService())
			if err != nil {
				// TODO
			}

			response, err := json.Marshal(users)
			if err != nil {
				// TODO
			}
			writer.WriteHeader(http.StatusOK)
			_, err = writer.Write(response)
			if err != nil {
				// TODO
			}
		})

		r.Patch("/user/{id}", func(writer http.ResponseWriter, request *http.Request) {
			errMessages := requests.UpdateUserValidate(request)
			if errMessages.Encode() != "" {
				err := badRequest(writer, errMessages)
				if err != nil {
					panic(err)
				}
				return
			}

			u, err := user.Update(request, provider.GetUserService())
			if err != nil {
				// TODO
			}

			response, err := json.Marshal(u)
			if err != nil {
				// TODO
			}
			writer.WriteHeader(http.StatusOK)
			_, err = writer.Write(response)
			if err != nil {
				// TODO
			}
		})

		r.Delete("/user/{id}", func(writer http.ResponseWriter, request *http.Request) {
			err := user.Delete(request, provider.GetUserService())
			if err != nil {
				// TODO
			}
			writer.WriteHeader(http.StatusNoContent)
		})
	})

	r.Group(func(r chi.Router) {
		r.Post("/post", func(writer http.ResponseWriter, request *http.Request) {
			errMessages := requests.StorePostValidate(request)

			if errMessages.Encode() != "" {
				err := badRequest(writer, errMessages)
				if err != nil {
					panic(err)
				}
				return
			}

			err := post.Store(request, provider.GetPostService())

			if err != nil {
				// TODO
			}

			response, err := json.Marshal("OK")
			if err != nil {
				// TODO
			}
			writer.WriteHeader(http.StatusCreated)
			_, err = writer.Write(response)
			if err != nil {
				// TODO
			}
		})

		r.Get("/posts", func(writer http.ResponseWriter, request *http.Request) {
			posts, err := post.GetAll(request, provider.GetPostService())
			if err != nil {
				// TODO
			}

			response, err := json.Marshal(posts)
			if err != nil {
				// TODO
			}

			writer.WriteHeader(http.StatusOK)
			_, err = writer.Write(response)
			if err != nil {
				// TODO
			}
		})
	})

	return r
}

func badRequest(writer http.ResponseWriter, errMessages url.Values) error {
	writer.WriteHeader(http.StatusBadRequest)

	response, err := json.Marshal(errMessages)
	if err != nil {
		return err
	}

	_, err = writer.Write(response)

	if err != nil {
		return err
	}

	return nil
}
