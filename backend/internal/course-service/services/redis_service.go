package services

import (
	"context"
	"github.com/doryngal/CourseFlow/backend/config"
	"github.com/go-redis/redis/v8"
	//"encoding/json"
)

type RedisService struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisService(cfg *config.CourseConfig) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return &RedisService{
		Client: rdb,
		Ctx:    context.Background(),
	}
}
