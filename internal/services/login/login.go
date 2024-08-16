package login

import (
	"VK-Pilot-Project/internal/models/auth"
	"context"
	"errors"

	"github.com/google/uuid"
)

func (service *Service) Login(ctx context.Context, auth auth.Model) (uuid.UUID, error) {
	user, err := service.repo.GetByLogin(ctx, auth.Login)
	if err != nil {
		return uuid.UUID{}, errors.New("no user with login " + auth.Login)
	}

	if !service.hasher.VerifuHash([]byte(user.Password), auth.Password) {
		return uuid.UUID{}, errors.New("wrong password")
	}
	return user.ID, nil
}
