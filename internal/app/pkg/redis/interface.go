package redis

import "context"

type RedisServiceInterface interface {
	Hset(ctx context.Context, key, field string, value interface{}) error
	HsetNX(ctx context.Context, key, field string, value interface{}) error
	Hget(ctx context.Context, key, field string, data interface{}) error
	HgetAll(ctx context.Context, key string) (map[string]string, error)
	SetNX(ctx context.Context, key string, data interface{}) error
	Get(ctx context.Context, key string, data interface{}) error
}
