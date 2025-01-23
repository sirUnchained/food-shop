package api

import (
	"foodshop/api/routers"
	"foodshop/configs"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *configs.Configs) {
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery())

	v1 := server.Group("/api/v1")
	initRoutes_v1(v1)

	server.Run(":4000")
}

func initRoutes_v1(route *gin.RouterGroup) {
	routers.UserRoutes(route)
}
