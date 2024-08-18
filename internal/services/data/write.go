package data

import (
	datamodel "VK-Pilot-Project/internal/models/data"
	"context"
)

func (service *Service) Write(ctx context.Context, write datamodel.Write) error {
	return service.repo.Write(ctx, write)
}
