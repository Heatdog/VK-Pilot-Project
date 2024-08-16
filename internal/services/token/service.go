package token

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Generate(ctx context.Context, id string) (token string, err error)
	Validate(ctx context.Context, token uuid.UUID) (fields TokenFields, err error)
}

type TokenFields struct {
	ID string `json:"sub"`
}
