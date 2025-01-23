package main

import (
	"fmt"
	"foodshop/configs"
)

func main() {
	fmt.Println("hello world !!")
	fmt.Println(configs.GetConfigs())
}
