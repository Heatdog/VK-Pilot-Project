package login

import (
	"VK-Pilot-Project/internal/repository/users"
	"VK-Pilot-Project/pkg/hash"
	"log/slog"
)

type Service struct {
	logger *slog.Logger
	repo   users.Repository
	hasher hash.Hasher
}

func New(logger *slog.Logger, repo users.Repository) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}
