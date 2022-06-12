package main

import (
	"chess-room/service"
	"chess-room/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func JudgeToken(ctx *gin.Context) {
	basicRequest := &service.BasicRequest{}
	basicRequest.Token = ctx.Request.Header.Get("token")
	response, err := roomClient.JudgeToken(ctx, basicRequest)
	if err != nil {
		fmt.Println("解析token失败", err)
	}
	tool.Info(ctx, response)
}
