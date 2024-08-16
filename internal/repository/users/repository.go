package users

import (
	"VK-Pilot-Project/internal/models/auth"
	"VK-Pilot-Project/internal/models/users"
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Insert(ctx context.Context, user auth.Model) (uuid.UUID, error)
	GetByLogin(ctx context.Context, login string) (users.Model, error)
}
