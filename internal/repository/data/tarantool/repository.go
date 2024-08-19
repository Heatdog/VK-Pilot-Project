package tarantool

import (
	datamodel "VK-Pilot-Project/internal/models/data"
	datarepo "VK-Pilot-Project/internal/repository/data"
	"context"
	"errors"
	"log/slog"

	tarantooldb "github.com/tarantool/go-tarantool/v2"
)

var _ datarepo.Repository = (*Repository)(nil)

type Repository struct {
	logger *slog.Logger
	conn   *tarantooldb.Connection
}

func New(logger *slog.Logger, conn *tarantooldb.Connection) *Repository {
	return &Repository{
		logger: logger,
		conn:   conn,
	}
}

const (
	spaceData = "data"
)

func (repo *Repository) Write(ctx context.Context, write datamodel.DataStruct) error {
	var futures []*tarantooldb.Future
	for key, val := range write.Data {
		request := tarantooldb.NewInsertRequest(spaceData).Tuple([]interface{}{
			key, val,
		})
		futures = append(futures, repo.conn.Do(request))
	}

	for _, future := range futures {
		result, err := future.Get()
		if err != nil {
			repo.logger.Error(err.Error())
			return err
		}
		repo.logger.Debug("insert", slog.Any("data", result))
	}
	return nil
}

func (repo *Repository) Read(ctx context.Context, read datamodel.KeysStruct) (datamodel.DataStruct, error) {
	var futures []*tarantooldb.Future
	for _, key := range read.Keys {
		request := tarantooldb.NewSelectRequest(spaceData).
			Iterator(tarantooldb.IterEq).
			Key([]interface{}{key})
		futures = append(futures, repo.conn.Do(request))
	}

	res := datamodel.DataStruct{
		Data: make(map[string]interface{}, len(read.Keys)),
	}

	for _, future := range futures {
		futResult, err := future.Get()
		if err != nil {
			repo.logger.Error(err.Error())
			return datamodel.DataStruct{}, err
		}

		repo.logger.Debug("get", slog.Any("tuple", futResult))

		if err = repo.parseData(res.Data, futResult); err != nil {
			repo.logger.Error(err.Error())
			return datamodel.DataStruct{}, err
		}
	}
	return res, nil
}

func (repo *Repository) parseData(res map[string]interface{}, data []interface{}) error {
	for _, el := range data {
		tuple, ok := el.([]interface{})
		if !ok {
			return errors.New("struct parse error")
		}

		if len(tuple) != 2 {
			return errors.New("struct parse error")
		}

		key, ok := tuple[0].(string)
		if !ok {
			return errors.New("key parse error")
		}

		res[key] = tuple[1]
	}

	return nil
}
