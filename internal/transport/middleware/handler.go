package middleware

import (
	"VK-Pilot-Project/internal/services/token"
	"log/slog"
)

type Handler struct {
	logger       *slog.Logger
	tokenService token.Service
}

func New(logger *slog.Logger, service token.Service) *Handler {
	return &Handler{
		logger:       logger,
		tokenService: service,
	}
}
