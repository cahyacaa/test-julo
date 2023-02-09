package redis

import (
	"context"
	"fmt"
	"github.com/cahyacaa/test-julo/cmd/config"
	"github.com/go-redis/redis/v8"
)

// nolint:revive
type RedisService struct{}

func NewRedisService() RedisService {
	return RedisService{}
}

var RedisInstance *redis.Client

func InitRedis(config config.Cache) error {
	RedisInstance = redis.NewClient(&redis.Options{
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
		Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
	})

	redisStatus := RedisInstance.Ping(context.Background())
	if redisStatus != nil {
		return redisStatus.Err()
	}
	return nil
}
