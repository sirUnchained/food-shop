package routers

import (
	"foodshop/api/controllers"
	middlewares "foodshop/api/middleWares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	handler := controllers.GetUserController()
	r.GET("/users", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), handler.GetAll)
	r.DELETE("/users/:id", middlewares.AuthorizeUser(), middlewares.RoleGaurd("admin"), handler.Remove)
	r.PUT("/users", middlewares.AuthorizeUser(), handler.Update)
	// r.PATCH("/users/:id/new-resturaunt", middlewares.AuthorizeUser(), handler.GetAll)
}
