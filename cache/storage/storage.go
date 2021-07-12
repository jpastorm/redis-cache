package storage

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

type storage struct {
	client *redis.Client
	context context.Context
}

func NewStorage(client *redis.Client) *storage {
	return &storage{
		client: client,
		context: context.Background(),
	}
}

// Create cache
func (s storage) Set(key string, value interface{}) error {
	err := s.client.Set(s.context, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s storage) Get(key string) (string, error) {
	result, err := s.client.Get(s.context, key).Result()
	if err == redis.Nil {
		return result, errors.New(key + " doesnt exists")
	}
	if err != nil {
		return result, err
	}
	return result, nil
}
// - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
func (s storage) HSet(key string, value map[string]interface{}) error {
	err := s.client.HMSet(s.context, key, value).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s storage) HGet(key ,field string) (string, error) {
	result, err := s.client.HGet(s.context, key, field).Result()
	if err == redis.Nil {
		return result, errors.New(key + " or " + field + " doesnt exists")
	}
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s storage) HMGet(key string , fields []string) ([]interface{}, error) {
	result, err := s.client.HMGet(s.context, key, fields...).Result()
	if err == redis.Nil {
		return result, errors.New(key + " or fields doesnt exists")
	}
	if err != nil {
		return result, err
	}
	return result, nil
}