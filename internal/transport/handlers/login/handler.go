package login

import (
	"VK-Pilot-Project/internal/services/login"
	"VK-Pilot-Project/internal/services/token"
	"VK-Pilot-Project/internal/transport/middleware"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	logger       *slog.Logger
	loginService *login.Service
	tokenService token.Service
	middleware   *middleware.Handler
}

const (
	loginURL = "/api/login"
)

func New(logger *slog.Logger, service *login.Service, mid *middleware.Handler, token token.Service) *handler {
	return &handler{
		logger:       logger,
		loginService: service,
		tokenService: token,
		middleware:   mid,
	}
}

func (handler *handler) HandleRoute(router *mux.Router) {
	router.HandleFunc(loginURL,
		handler.middleware.Recover(
			handler.middleware.Logging(handler.login))).
		Methods(http.MethodPost)
}
