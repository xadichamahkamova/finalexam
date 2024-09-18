package redis

import (
	"context"
	"fmt"
	config "api-gateway/internal/pkg/load"

	"github.com/redis/go-redis/v9"
)

type Redis struct{
	Client redis.Client
}

func NewClient(cfg config.Config) (*Redis, error){

	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",
		DB: 0,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &Redis{
		Client: *rdb,
	}, nil
}