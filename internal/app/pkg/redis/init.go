package redis

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/cahyacaa/test-julo/cmd/config"
	"github.com/redis/go-redis/v9"
)

// nolint:revive
type RedisService struct {
	redisDB   *redis.Client
	RedisLock *redislock.Client
}

func NewRedisService() RedisService {
	return RedisService{}
}

func (r *RedisService) InitRedis(config config.Cache) error {
	redisOpts := redis.Options{
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
		Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
	}
	r.redisDB = redis.NewClient(&redisOpts)

	redisStatus := r.redisDB.Ping(context.Background()).Err()
	if redisStatus != nil {
		return redisStatus
	}

	r.RedisLock = redislock.New(r.redisDB)
	return nil
}
