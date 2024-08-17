package token

import "context"

type Service interface {
	Generate(ctx context.Context, id string) (token string, err error)
	Validate(ctx context.Context, token string) (fields TokenFields, err error)
}

type TokenFields struct {
	ID string `json:"sub"`
}
