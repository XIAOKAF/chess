package tool

import (
	"chess/proto"
	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context, response *proto.Response) {
	ctx.JSON(int(response.Code), gin.H{
		"code": response.Code,
		"info": response.Info,
	})
}
