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

func Join(ctx *gin.Context) {
	joinRequest := &service.JoinRequest{}
	joinRequest.Token = ctx.Request.Header.Get("token")
	joinRequest.RoomId = ctx.PostForm("roomId")
	response, err := roomClient.Join(ctx, joinRequest)
	if err != nil {
		fmt.Println("加入游戏房间错误", err)
	}
	tool.Info(ctx, response)
}

func UpdateStatus(ctx *gin.Context) {
	updateRequest := &service.UpdateRequest{}
	updateRequest.Token = ctx.Request.Header.Get("token")
	updateRequest.RoomId = ctx.PostForm("roomId")
	updateRequest.Status = ctx.PostForm("status")
	response, err := roomClient.UpdateStatus(ctx, updateRequest)
	if err != nil {
		fmt.Println("更新状态失败", err)
	}
	tool.Info(ctx, response)
}
