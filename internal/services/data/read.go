package data

import (
	datamodel "VK-Pilot-Project/internal/models/data"
	"context"
)

func (service *Service) Read(ctx context.Context, keys datamodel.KeysStruct) (datamodel.DataStruct, error) {
	return service.repo.Read(ctx, keys)
}
