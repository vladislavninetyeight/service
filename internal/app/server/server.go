package server

import (
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/app/router"
	"github.com/vladislavninetyeight/service/tree/main/internal/model/internal/providers"
	"net/http"
)

type Server struct {
	http *http.Server
}

func NewServer(provider *providers.ServiceProvider) *Server {
	return &Server{
		http: &http.Server{
			Addr:    ":8080",
			Handler: router.New(*provider),
		},
	}
}

func (s *Server) Run() {

	err := s.http.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
