package jwt

import (
	servicetoken "VK-Pilot-Project/internal/services/token"
	"context"
	"fmt"
	"time"

	jwtlib "github.com/golang-jwt/jwt"
)

var _ servicetoken.Service = (*Service)(nil)

type Service struct {
	key string
}

func New(key string) *Service {
	return &Service{
		key: key,
	}
}

func (service *Service) Generate(ctx context.Context, id string) (string, error) {
	payload := jwtlib.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, payload)
	return token.SignedString(service.key)
}

func (service *Service) Validate(ctx context.Context, token string) (servicetoken.TokenFields, error) {
	parsedToken, err := jwtlib.Parse(token, func(t *jwtlib.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(service.key), nil
	})

	if err != nil {
		return servicetoken.TokenFields{}, err
	}

	claims, ok := parsedToken.Claims.(jwtlib.MapClaims)
	if !ok {
		return servicetoken.TokenFields{}, fmt.Errorf("token calims are not of type *TokenClaims")
	}

	res := servicetoken.TokenFields{
		ID: claims["sub"].(string),
	}
	return res, nil

}
