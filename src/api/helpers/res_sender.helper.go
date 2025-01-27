package helpers

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

type ResultResponse struct {
	Ok      bool        `json:"ok"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResultResponse(ok bool, status int, message string, data interface{}) *ResultResponse {
	return &ResultResponse{
		Ok:      ok,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func SendResult(ok bool, status int, msg string, data interface{}, ctx *gin.Context) {
	var res ResultResponse
	res.Data = data
	res.Message = msg
	res.Ok = ok
	res.Status = status

	ctx.JSON(status, res)
}

func SendValidationErrors(status int, errs string, ctx *gin.Context) {
	var res ResultResponse
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

func SendUnAuthorizedResult(ctx *gin.Context) {
	var res ResultResponse

	res.Ok = false
	res.Status = 401
	res.Message = "please sign in or sign up first."
	res.Data = nil

	ctx.JSON(401, res)
}
