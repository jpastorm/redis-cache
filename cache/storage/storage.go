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