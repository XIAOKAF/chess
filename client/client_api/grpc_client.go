package main

import (
	"chess-room/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var roomClient service.RoomServiceClient

func main() {
	engine := gin.Default()
	conn, err := grpc.Dial(":8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("连接注册中心失败", err)
	}
	roomClient = service.NewRoomServiceClient(conn)

	engine.Use(JudgeToken)
	engine.POST("/create", Create)             //创建游戏房间
	engine.POST("/join", Join)                 //加入游戏房间
	engine.POST("/updateStatus", UpdateStatus) //更新玩家状态
	engine.DELETE("/exit")

	engine.Run(":/8001")
}
