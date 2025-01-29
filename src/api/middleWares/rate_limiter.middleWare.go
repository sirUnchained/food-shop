package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth/v8"
	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)

	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "too many request !"})
			return
		} else {
			ctx.Next()
		}
	}
}
