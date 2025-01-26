package routers

import (
	"foodshop/api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	controller := controllers.GetAuthController()

	r.POST("/auth/login", controller.Login)
	r.POST("/auth/register", controller.Register)
}
