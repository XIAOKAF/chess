package main

import (
	"chess-room/service"
	"chess-room/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	createRequest := &service.CreateRequest{}
	createRequest.Token = ctx.Request.Header.Get("token")
	response, err := roomClient.Create(ctx, createRequest)
	if err != nil {
		fmt.Println("创建游戏房间失败", err)
	}
	tool.Info(ctx, response)
}
