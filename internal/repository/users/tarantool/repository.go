package tarantool

import (
	"VK-Pilot-Project/internal/models/auth"
	usersmodel "VK-Pilot-Project/internal/models/users"
	"VK-Pilot-Project/internal/repository/users"
	"VK-Pilot-Project/pkg/hash"
	"context"
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
	loginIndx  = "login"
)

func New(logger *slog.Logger, conn *tarantooldb.Connection) (*Repository, error) {
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
		Tuple([]interface{}{id.String(), user.Login, string(hashedPSWD)}))

	if _, err := future.Get(); err != nil {
		return uuid.UUID{}, err
	}

	repo.logger.Debug("insert", slog.String("id", id.String()), slog.String("login", user.Login))
	return id, nil
}

func (repo *Repository) GetByLogin(ctx context.Context, login string) (usersmodel.Model, bool) {
	res, err := repo.conn.Do(tarantooldb.NewSelectRequest(spaceUsers).
		Index(loginIndx).
		Iterator(tarantooldb.IterEq).
		Key([]interface{}{login})).
		Get()

	if err != nil {
		repo.logger.Error(err.Error())
		return usersmodel.Model{}, false
	}

	repo.logger.Debug("results", slog.Any("users", res))

	el, ok := repo.parseUser(res, login)
	if !ok {
		return usersmodel.Model{}, false
	}

	repo.logger.Debug("get", slog.Any("user", el))

	return el, true
}

func (reo *Repository) parseUser(data []interface{}, login string) (usersmodel.Model, bool) {
	if len(data) != 1 {
		return usersmodel.Model{}, false
	}

	tuple, ok := data[0].([]interface{})
	if !ok {
		return usersmodel.Model{}, false
	}

	id, ok := tuple[0].(string)
	if !ok {
		return usersmodel.Model{}, false
	}

	hashedPSWD, ok := tuple[2].(string)
	if !ok {
		return usersmodel.Model{}, false
	}

	el := usersmodel.Model{
		ID:       id,
		Login:    login,
		Password: hashedPSWD,
	}
	return el, true
}
