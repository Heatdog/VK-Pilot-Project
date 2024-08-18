package login

import (
	"VK-Pilot-Project/internal/models/auth"
	"context"
	"errors"
)

func (service *Service) Login(ctx context.Context, auth auth.ModelRequest) (string, error) {
	user, ok := service.repo.GetByLogin(ctx, auth.Login)
	if !ok {
		return "", errors.New("no user with login " + auth.Login)
	}

	if !service.hasher.VerifyHash([]byte(user.Password), auth.Password) {
		return "", errors.New("wrong password")
	}
	return user.ID, nil
}
