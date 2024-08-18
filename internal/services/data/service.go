package data

import (
	datarepo "VK-Pilot-Project/internal/repository/data"
	"log/slog"
)

type Service struct {
	logger *slog.Logger
	repo   datarepo.Repository
}

func New(logger *slog.Logger, repo datarepo.Repository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}
