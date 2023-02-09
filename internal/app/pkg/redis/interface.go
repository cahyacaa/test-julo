package redis

import "context"

//go:generate mockgen -destination=mocks/mock.go -source=interface.go RepositoryInterface

type RedisServiceInterface interface {
	Hset(ctx context.Context, key, field string, value interface{}) error
	HsetNX(ctx context.Context, key, field string, value interface{}) error
	Hget(ctx context.Context, key, field string) (string, error)
	SetNX(ctx context.Context, key string, data interface{}) error
	Exists(ctx context.Context, key string) bool
	Get(ctx context.Context, key string, data interface{}) error
}
