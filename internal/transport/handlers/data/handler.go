package data

import (
	dataservice "VK-Pilot-Project/internal/services/data"
	"VK-Pilot-Project/internal/transport/middleware"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	logger     *slog.Logger
	service    *dataservice.Service
	middleware *middleware.Handler
}

const (
	loginURL = "/api/write"
)

func New(logger *slog.Logger, service *dataservice.Service, mid *middleware.Handler) *handler {
	return &handler{
		logger:     logger,
		service:    service,
		middleware: mid,
	}
}

func (handler *handler) HandleRoute(router *mux.Router) {
	router.HandleFunc(loginURL,
		handler.middleware.Recover(
			handler.middleware.Logging(
				handler.middleware.Auth(handler.write)))).
		Methods(http.MethodPost)
}
