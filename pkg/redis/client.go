package redis

import (
	"context"
	"fmt"

	internal "github.com/Qwepo/InCryipt/internal"

	"github.com/redis/go-redis/v9"
)

func NewClient(ctx context.Context, db int, conf *internal.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", conf.Redis.Address, conf.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Redis.Password,
		DB:       db,
	})

	defer client.Close()

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
