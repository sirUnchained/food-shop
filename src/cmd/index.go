package main

import (
	"foodshop/api"
	"foodshop/configs"
	"foodshop/data/postgres"
)

func main() {
	cfg := configs.GetConfigs()
	api.InitServer(cfg)

	// config postgres
	err := postgres.InitPostgres(cfg)
	if err != nil {
		panic(err)
	}
	defer postgres.CloseDb()
}
