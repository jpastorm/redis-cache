package usecase

import "github.com/jpastorm/redis-cache/cache/model"

type Usecase interface {
	Set(key string, value model.Cache) error
	Get(key string) (string, error)
}

type Storage interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
	HSet(key string, value map[string]interface{}) error
	HGet(key ,field string) (string, error)
	HMGet(key string , fields []string) ([]interface{}, error)
}