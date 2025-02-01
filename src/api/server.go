package api

import (
	middlewares "foodshop/api/middleWares"
	"foodshop/api/routers"
	validators "foodshop/api/validations"
	"foodshop/configs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *configs.Configs) {
	server := gin.New()
	server.Use(gin.Logger(), gin.Recovery(), middlewares.Limiter())

	InitValidators()

	v1 := server.Group("/api/v1")
	initRoutes_v1(v1)

	server.Run(cfg.Server.Port)
}

func initRoutes_v1(route *gin.RouterGroup) {
	routers.UserRoutes(route)
	routers.AuthRoutes(route)
	routers.CategoryRoutes(route)
	routers.RestaurantRoutes(route)
	routers.FoodRoutes(route)
}

func InitValidators() {
	validation, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		validation.RegisterValidation("IranMobile", validators.IranMobileValidator)
	}
}
