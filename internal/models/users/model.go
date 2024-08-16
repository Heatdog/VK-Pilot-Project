package users

import "github.com/google/uuid"

type Model struct {
	ID       uuid.UUID
	Login    string
	Password string
}
