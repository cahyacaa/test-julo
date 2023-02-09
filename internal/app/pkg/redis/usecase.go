package redis

import (
	"context"
	"encoding/json"
	"fmt"
)

// SetNX a key/value
func (r *RedisService) SetNX(ctx context.Context, key string, data interface{}) error {

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	result := r.redisDB.SetNX(ctx, key, value, 0)

	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

func (r *RedisService) Set(ctx context.Context, key string, data interface{}) error {

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	result := r.redisDB.Set(ctx, key, value, 0)

	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

// Exists check a key
func (r *RedisService) Exists(ctx context.Context, key string) bool {

	result, err := r.redisDB.Exists(ctx, key).Result()

	if result == 0 || err != nil {
		return false
	}

	return true
}

// Get get a key
func (r *RedisService) Get(ctx context.Context, key string, data interface{}) error {

	result, err := r.redisDB.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisService) Hset(ctx context.Context, key, field string, value interface{}) error {
	value, err := json.Marshal(value)
	if err != nil {
		return err
	}

	a, err := r.redisDB.HSet(ctx, key, field, value).Result()
	if err != nil {
		return err
	}

	fmt.Println(a)

	return nil
}

func (r *RedisService) HsetNX(ctx context.Context, key, field string, value interface{}) error {

	value, err := json.Marshal(value)
	if err != nil {
		return err
	}

	_, err = r.redisDB.HSetNX(ctx, key, field, value).Result()

	if err != nil {
		return err
	}

	return nil
}

func (r *RedisService) Hget(ctx context.Context, key, field string, data interface{}) error {
	result, err := r.redisDB.HGet(ctx, key, field).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisService) HgetAll(ctx context.Context, key string) (map[string]string, error) {
	result, err := r.redisDB.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]string{}, err
	}

	return result, nil
}
