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
	writeURL = "/api/write"
	readURL  = "/api/read"
)

func New(logger *slog.Logger, service *dataservice.Service, mid *middleware.Handler) *handler {
	return &handler{
		logger:     logger,
		service:    service,
		middleware: mid,
	}
}

func (handler *handler) HandleRoute(router *mux.Router) {
	router.HandleFunc(writeURL,
		handler.middleware.Recover(
			handler.middleware.Logging(
				handler.middleware.Auth(handler.write)))).
		Methods(http.MethodPost)

	router.HandleFunc(readURL,
		handler.middleware.Recover(
			handler.middleware.Logging(
				handler.middleware.Auth(handler.read)))).
		Methods(http.MethodPost)
}
