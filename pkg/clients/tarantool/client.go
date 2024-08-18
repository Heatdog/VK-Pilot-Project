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
		Address: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
	}

	opt := tarantooldb.Opts{
		Timeout: time.Duration(cfg.TimeoutSeconds) * time.Second,
	}

	conn, err := tarantooldb.Connect(ctx, dialer, opt)
	if err != nil {
		return nil, err
	}

	if _, err = conn.Do(tarantooldb.NewPingRequest()).Get(); err != nil {
		return nil, err
	}

	return conn, nil
}
