package data

import (
	"VK-Pilot-Project/internal/transport/middleware"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	logger     *slog.Logger
	middleware *middleware.Handler
}

const (
	loginURL = "/api/login"
)

func New(logger *slog.Logger, mid *middleware.Handler) *handler {
	return &handler{
		logger:     logger,
		middleware: mid,
	}
}

func (handler *handler) HandleRoute(router *mux.Router) {
	router.HandleFunc(loginURL,
		handler.middleware.Recover(
			handler.middleware.Logging(handler.write))).
		Methods(http.MethodPost)
}
