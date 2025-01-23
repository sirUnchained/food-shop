package api

import (
	"foodshop/configs"

	"github.com/gin-gonic/gin"
)

func InitServer(cfg *configs.Configs) {
	server := gin.New()
	v1 := server.Group("/api/v1")
	initRoutes(v1)

	server.Run(":4000")
}

// initialize routes
func initRoutes(route *gin.RouterGroup) {

}
