package redis

import (
	"context"
	"encoding/json"
)

// Set a key/value
func (r RedisService) Set(ctx context.Context, key string, data interface{}) error {

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	result := RedisInstance.Set(ctx, key, value, 0)

	if result.Err() != nil {
		return result.Err()
	}

	return nil
}

// Exists check a key
func (r RedisService) Exists(ctx context.Context, key string) bool {

	result, err := RedisInstance.Exists(ctx, key).Result()

	if result == 0 || err != nil {
		return false
	}

	return true
}

// Get get a key
func (r RedisService) Get(ctx context.Context, key string, data interface{}) error {

	result, err := RedisInstance.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(result), &data)
	if err != nil {
		return err
	}

	return nil
}

// Delete delete a key
func (r RedisService) Delete(ctx context.Context, key ...string) (bool, error) {

	result, err := RedisInstance.Del(ctx, key...).Result()

	if err != nil || result == 0 {
		return false, nil
	}
	return true, nil
}
