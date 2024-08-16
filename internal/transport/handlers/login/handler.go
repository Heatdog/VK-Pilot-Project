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
}

const (
	loginURL = "/login"
)

func HandleRoute(router *mux.Router, logger *slog.Logger, loginService *login.Service,
	mid *middleware.Handler, tokenService token.Service) {
	handler := &handler{
		logger:       logger,
		loginService: loginService,
		tokenService: tokenService,
	}

	router.HandleFunc(loginURL, mid.Recover(mid.Logging(handler.login))).Methods(http.MethodPost)
}
