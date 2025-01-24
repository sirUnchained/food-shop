package helpers

import (
	"github.com/gin-gonic/gin"
)

type ResultResponse struct {
	Ok      bool        `json:"ok"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// type ValidationFaliedResponse struct

func SendResult(ok bool, status int, msg string, data *interface{}, ctx *gin.Context) {
	var res ResultResponse
	res.Data = data
	res.Message = msg
	res.Ok = ok
	res.Status = status

	ctx.JSON(status, res)
	return
}
