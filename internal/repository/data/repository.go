package data

import (
	datamodel "VK-Pilot-Project/internal/models/data"
	"context"
)

type Repository interface {
	Write(ctx context.Context, write datamodel.Write) error
	Read(ctx context.Context, read datamodel.Read) (datamodel.Write, error)
}
