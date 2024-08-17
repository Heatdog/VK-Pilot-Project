package tarantool

import (
	"VK-Pilot-Project/internal/models/auth"
	usersmodel "VK-Pilot-Project/internal/models/users"
	"VK-Pilot-Project/internal/repository/users"
	"VK-Pilot-Project/pkg/hash"
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	tarantooldb "github.com/tarantool/go-tarantool/v2"
)

var _ users.Repository = (*Repository)(nil)

type Repository struct {
	logger *slog.Logger
	conn   *tarantooldb.Connection
	hasher hash.Hasher
}

const (
	spaceUsers = "users"
)

func New(logger *slog.Logger, conn *tarantooldb.Connection) (*Repository, error) {
	future := conn.Do(tarantooldb.NewCallRequest("box.schema.space.create").
		Args([]interface{}{
			spaceUsers,
			map[string]bool{"if_not_exists": true},
		}))

	if _, err := future.Get(); err != nil {
		return nil, err
	}
	/*

		future = conn.Do(tarantooldb.NewCallRequest("box.space.users:format").
			Args([][]map[string]string{
				{
					{
						"name": "id",
						"type": "uuid",
					},
					{
						"name": "login",
						"type": "string",
					},
					{
						"name": "password",
						"type": "string",
					},
				},
			}))

		if _, err := future.Get(); err != nil {
			return nil, err
		}

		future = conn.Do(tarantooldb.NewCallRequest("box.space.users:create_index").
			Args([]interface{}{
				"primary",
				map[string]interface{}{
					"parts":         []string{"id"},
					"if_not_exists": true,
				},
			}))

		if _, err := future.Get(); err != nil {
			return nil, err
		}

		future = conn.Do(tarantooldb.NewCallRequest("box.space.users:create_index").
			Args([]interface{}{
				"loginidx",
				map[string]interface{}{
					"parts":         []string{"login"},
					"type":          "BTREE",
					"unique":        true,
					"if_not_exists": true,
				},
			}))

		if _, err := future.Get(); err != nil {
			return nil, err
		}
	*/
	return &Repository{
		logger: logger,
		conn:   conn,
	}, nil
}

func (repo *Repository) Insert(ctx context.Context, user auth.ModelRequest) (uuid.UUID, error) {
	id := uuid.New()

	hashedPSWD, err := repo.hasher.Hash(user.Password)
	if err != nil {
		return uuid.UUID{}, err
	}

	future := repo.conn.Do(tarantooldb.
		NewInsertRequest(spaceUsers).
		Tuple([]usersmodel.Model{{
			ID:       id,
			Login:    user.Login,
			Password: string(hashedPSWD),
		}}))

	if _, err := future.Get(); err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (repo *Repository) GetByLogin(ctx context.Context, login string) (usersmodel.Model, error) {
	data, err := repo.conn.Do(tarantooldb.NewSelectRequest(spaceUsers).
		Index("login").
		Key([]interface{}{login})).
		Get()

	if err != nil {
		return usersmodel.Model{}, err
	}

	res, ok := data[0].(usersmodel.Model)
	if !ok {
		return usersmodel.Model{}, errors.New("bad cast")
	}

	return res, nil
}
