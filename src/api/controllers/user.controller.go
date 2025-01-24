package controllers

import (
	"fmt"
	"foodshop/api/helpers"
	"foodshop/data/redis"
	"time"

	"github.com/gin-gonic/gin"
)

type userController struct{}

func GetUserController() *userController {
	return &userController{}
}

func (u *userController) GetAll(ctx *gin.Context) {
	err := redis.SetInRedis(redis.GetRedisClient(), ctx, "test", "users", time.Second*60)
	if err != nil {
		fmt.Println(err)
	}

	result, err := redis.GetFromRedis[string](redis.GetRedisClient(), ctx, "test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	helpers.SendResult(true, 200, "success", nil, ctx)
}
