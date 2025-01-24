package main

import (
	"fmt"
	"foodshop/api"
	"foodshop/configs"
)

func main() {
	fmt.Println("hello world !!")
	cfg := configs.GetConfigs()
	api.InitServer(cfg)
	// fmt.Println(configs.GetConfigs())
}
