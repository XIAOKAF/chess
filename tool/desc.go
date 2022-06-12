package tool

import (
	"chess-room/service"
	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context, response *service.Response) {
	ctx.JSON(int(response.Code), gin.H{
		"code": response.Code,
		"info": response.Info,
	})
}

func Failure(response *service.Response, code int32, info string) *service.Response {
	response.Code = code
	response.Info = info
	return response
}

func Success(response *service.Response, code int32, info string) *service.Response {
	response.Code = code
	response.Info = info
	return response
}
