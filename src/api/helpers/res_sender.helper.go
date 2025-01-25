package helpers

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

type resultResponse struct {
	Ok      bool        `json:"ok"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResult(ok bool, status int, msg string, data *interface{}, ctx *gin.Context) {
	var res resultResponse
	res.Data = data
	res.Message = msg
	res.Ok = ok
	res.Status = status

	ctx.JSON(status, res)
}

func SendValidationErrors(status int, errs string, ctx *gin.Context) {
	var res resultResponse
	pattern := regexp.MustCompile(`Error:.*?tag`)

	validationErrors := pattern.FindAllString(errs, -1)
	if len(validationErrors) == 0 {
		validationErrors = append(validationErrors, errs)
	}

	res.Message = "Validation Failed"
	res.Ok = false
	res.Status = status
	res.Data = validationErrors

	ctx.JSON(status, res)
}
