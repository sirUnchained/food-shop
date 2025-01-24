package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"foodshop/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *configs.Configs) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		DB:   cfg.Redis.Db,
	})

	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedisClient() {
	redisClient.Close()
}

func SetInRedis[T any](rc *redis.Client, ctx *gin.Context, key string, value T, expiration time.Duration) error {
	str, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = rc.Set(ctx, key, string(str), expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetFromRedis[T any](rc *redis.Client, ctx *gin.Context, key string) (T, error) {
	var dest T = *new(T)

	result, err := rc.Get(ctx, key).Result()
	if err != nil {
		return dest, err
	}

	err = json.Unmarshal([]byte(result), &dest)
	if err != nil {
		return dest, err
	}

	return dest, nil
}
