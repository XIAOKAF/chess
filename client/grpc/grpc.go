package main

import (
	"chess/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var userClient proto.UserServiceClient

func main() {
	engine := gin.Default()
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("连接用户中心失败", err)
	}
	//注册所有服务
	userClient = proto.NewUserServiceClient(conn)
	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", Register)
		userGroup.POST("/login", Login)
	}
	smsGroup := engine.Group("/sms")
	{
		smsGroup.POST("/send", Send)
	}
	engine.Run(":8000")
}
