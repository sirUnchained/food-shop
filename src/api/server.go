package api

import (
	middlewares "foodshop/api/middleWares"
	"foodshop/api/routers"
	"foodshop/configs"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *configs.Configs) {
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery(), middlewares.Limiter())

	v1 := server.Group("/api/v1")
	initRoutes_v1(v1)

	server.Run(cfg.Server.Port)
}

func initRoutes_v1(route *gin.RouterGroup) {
	// todo => before coding, follow this shit and eplain to your self.
	routers.UserRoutes(route)
}
