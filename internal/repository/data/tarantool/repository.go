package tarantool

import (
	datamodel "VK-Pilot-Project/internal/models/data"
	datarepo "VK-Pilot-Project/internal/repository/data"
	"context"
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

func (repo *Repository) Write(ctx context.Context, write datamodel.Write) error {
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

func (repo *Repository) Read(ctx context.Context, read datamodel.Read) (datamodel.Write, error) {
	return datamodel.Write{}, nil
}
