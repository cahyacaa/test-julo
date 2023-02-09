package redis

import "context"

//go:generate mockgen -destination=mocks/mock.go -source=interface.go RepositoryInterface

type RedisServiceInterface interface {
	SetNX(ctx context.Context, key string, data interface{}) error
	Exists(ctx context.Context, key string) bool
	Get(ctx context.Context, key string, data interface{}) error
	Incr(ctx context.Context, key string) (int64, error)
	Delete(ctx context.Context, key ...string) (bool, error)
}
