package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/junstory/go_todo_app/week8/config"
	"github.com/junstory/go_todo_app/week8/entity"
)

type KVS struct {
	Cli *redis.Client
}

func NewKVS(ctx context.Context, cfg *config.Config) (*KVS, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &KVS{Cli: cli}, nil
}

func (k *KVS) Save(ctv context.Context, key string, userID entity.UserID) error {
	id := int64(userID)
	return k.Cli.Set(ctv, key, id, 30*time.Minute).Err()
}

func (k *KVS) Load(ctv context.Context, key string) (entity.UserID, error) {
	id, err := k.Cli.Get(ctv, key).Int64()
	if err != nil {
		return 0, fmt.Errorf("failed to get key %q: %w", key, ErrNotFound)
	}
	return entity.UserID(id), nil
}
