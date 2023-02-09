package redis

import "context"

//go:generate mockgen -destination=mocks/mock.go -source=interface.go RepositoryInterface

type RedisServiceInterface interface {
	Set(ctx context.Context, key string, data interface{}) error
	Exists(ctx context.Context, key string) bool
	Get(ctx context.Context, key string, data interface{}) error
	Delete(ctx context.Context, key ...string) (bool, error)
}
