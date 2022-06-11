package main

import (
	"chess-room/server/controller"
	"chess-room/server/dao"
	"chess-room/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	dao.InitRDB()
	dao.InitMDB()
	dao.InitDB()
	rpcServer := grpc.NewServer()
	service.RegisterRoomServiceServer(rpcServer, controller.RoomService)
	listener, err := net.Listen("tcp", ":8003")
	if err != nil {
		log.Fatal("启动监听失败", err)
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动服务失败", err)
	}
}
