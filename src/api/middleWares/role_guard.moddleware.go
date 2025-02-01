package middlewares

import (
	"foodshop/api/helpers"
	"foodshop/data/models"

	"github.com/gin-gonic/gin"
)

func RoleGaurd(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _ := ctx.Get("user")
		roles := user.(models.Users).Roles

		for _, r := range roles {
			if r.State == role {
				ctx.Next()
				return
			}
		}

		helpers.SendResult(false, 403, "this route is protected and you can't have access.", nil, ctx)
		ctx.Abort()
		return
	}
}
