package tarantool

import (
	"VK-Pilot-Project/internal/models/auth"
	"VK-Pilot-Project/internal/repository/users/tarantool"
	"context"
)

func Init(ctx context.Context, repo *tarantool.Repository) error {
	users := []auth.ModelRequest{
		{
			Login:    "admin",
			Password: "presale",
		},
	}

	for _, user := range users {
		if _, err := repo.Insert(ctx, user); err != nil {
			return err
		}
	}
	return nil
}
