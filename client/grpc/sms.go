package main

import (
	"chess/client/tool"
	"chess/proto"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Send(ctx *gin.Context) {
	sendRequest := &proto.SendRequest{}
	sendRequest.Mobile = ctx.PostForm("mobile")
	response, err := userClient.Send(ctx, sendRequest)
	if err != nil {
		fmt.Println("发送短信失败", err)
	}
	tool.Info(ctx, response)
}
