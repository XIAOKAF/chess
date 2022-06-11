package tool

import (
	"chess/service"
	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context, response *service.Response) {
	ctx.JSON(int(response.Code), gin.H{
		"code": response.Code,
		"info": response.Info,
	})
}
