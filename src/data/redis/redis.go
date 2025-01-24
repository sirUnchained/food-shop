package redis

import (
	"context"
	"fmt"
	"foodshop/configs"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *configs.Configs) error {
	red := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		DB:   cfg.Redis.Db,
	})

	ctx := context.Background()
	_, err := red.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	if redisClient != nil {
		redisClient.Close()
	}
}
