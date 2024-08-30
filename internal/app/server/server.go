package server

import (
	"fmt"
	"github.com/vladislavninetyeight/service/internal/config"
	"net/http"
)

func New(config *config.Server) *http.Server {
	address := fmt.Sprintf(":%s", config.Port)

	return &http.Server{
		Addr: address,
	}
}
