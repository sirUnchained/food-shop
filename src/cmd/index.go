package main

import (
	"foodshop/api"
	"foodshop/configs"
	"foodshop/data/postgres"
	"foodshop/data/redis"
)

func main() {
	cfg := configs.GetConfigs()

	// config postgres
	err := postgres.InitPostgres(cfg)
	if err != nil {
		panic(err)
	}
	defer postgres.CloseDb()

	// config redis
	err = redis.InitRedis(cfg)
	if err != nil {
		panic(err)
	}
	defer redis.CloseRedisClient()

	api.InitServer(cfg)
}
