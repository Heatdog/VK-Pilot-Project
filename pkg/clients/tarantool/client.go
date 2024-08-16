package tarantool

import (
	"VK-Pilot-Project/internal/config"
	"context"
	"fmt"
	"time"

	tarantooldb "github.com/tarantool/go-tarantool/v2"
)

func NewClient(ctx context.Context, cfg config.Tarantool) (*tarantooldb.Connection, error) {
	dialer := tarantooldb.NetDialer{
		Address: fmt.Sprintf("%s:%d", cfg.IP, cfg.Port),
		User:    cfg.User,
	}

	opt := tarantooldb.Opts{
		Timeout: time.Duration(cfg.TimeoutSeconds) * time.Second,
	}

	return tarantooldb.Connect(ctx, dialer, opt)
}
