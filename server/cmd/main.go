package main

import (
	"chess/proto"
	"chess/server/controller"
	"chess/server/dao"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	dao.InitRDB()
	dao.InitMDB()
	dao.InitDB()
	rpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(rpcServer, controller.UserService)
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听失败", err)
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("启动服务失败", err)
	}
}
