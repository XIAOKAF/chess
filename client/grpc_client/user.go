package main

import (
	"chess/client/tool"
	"chess/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	registerRequest := &service.RegisterRequest{}
	registerRequest.Mobile = ctx.PostForm("mobile")
	registerRequest.Password = ctx.PostForm("pwd")
	registerRequest.ConfirmPwd = ctx.PostForm("confirmPwd")
	registerRequest.Code = ctx.PostForm("code")
	response, err := userClient.Register(ctx, registerRequest)
	if err != nil {
		fmt.Println("注册失败", err)
	}
	tool.Info(ctx, response)
}

func Login(ctx *gin.Context) {
	loginRequest := &service.LoginRequest{}
	loginRequest.Mobile = ctx.PostForm("mobile")
	loginRequest.Password = ctx.PostForm("pwd")
	response, err := userClient.Login(ctx, loginRequest)
	if err != nil {
		fmt.Println("登录失败", err)
	}
	tool.Info(ctx, response)
}
