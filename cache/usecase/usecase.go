package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/jpastorm/redis-cache/cache/model"
)

type Cache struct {
	storage Storage
}

func NewCache(storage Storage) Cache{
	return Cache{storage: storage}
}

// Create creates a new model.User
func (c Cache) Set(key string, value model.Cache) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := c.storage.Set(key, bytes); err != nil {
		return fmt.Errorf("set: %w", err)
	}

	return nil
}

func (c Cache) Get(key string) (string, error) {
	result, err := c.storage.Get(key)
	if err != nil {
		return "",fmt.Errorf("get: %w", err)
	}
	return result, nil
}