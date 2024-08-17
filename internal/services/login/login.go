package login

import (
	"VK-Pilot-Project/internal/models/auth"
	"context"
	"errors"
)

func (service *Service) Login(ctx context.Context, auth auth.ModelRequest) (string, error) {
	user, err := service.repo.GetByLogin(ctx, auth.Login)
	if err != nil {
		return "", errors.New("no user with login " + auth.Login)
	}

	if !service.hasher.VerifuHash([]byte(user.Password), auth.Password) {
		return "", errors.New("wrong password")
	}
	return user.ID, nil
}
