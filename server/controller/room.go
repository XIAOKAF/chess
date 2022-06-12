package controller

import (
	"chess-room/server/model"
	"chess-room/server/serverService"
	"chess-room/service"
	"chess-room/tool"
	"context"
)

var RoomService = &roomService{}

type roomService struct {
}

func (r *roomService) Create(ctx context.Context, request *service.CreateRequest) (*service.Response, error) {
	resp := &service.Response{}
	tokenClaims, _ := serverService.ParseToken(request.Token)
	room := &model.Room{
		Owner: tokenClaims.Mobile,
	}
	err := serverService.InsertRoom(room)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, nil
	}
	err = serverService.RecordRoom(room)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, nil
	}
	id, err := serverService.SelectRoom(room)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, nil
	}
	resp = tool.Success(resp, 200, id)
	return resp, nil
}

func (r *roomService) Join(ctx context.Context, request *service.JoinRequest) (*service.Response, error) {
	resp := &service.Response{}
	tokenClaims, _ := serverService.ParseToken(request.Token)

	player := &model.Player{
		Mobile: tokenClaims.Mobile,
		RoomId: request.RoomId,
	}
	err, flag := serverService.GetPlayer(player, "two")
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	if flag {
		response := tool.Failure(resp, 400, "该房间人数已满")
		return response, nil
	}
	err = serverService.Join(player)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	resp = tool.Success(resp, 200, "成功加入房间")
	return resp, nil
}

func (r *roomService) UpdateStatus(ctx context.Context, request *service.UpdateRequest) (*service.Response, error) {
	resp := &service.Response{}
	tokenClaims, _ := serverService.ParseToken(request.Token)
	player := &model.Player{
		Mobile: tokenClaims.Mobile,
		RoomId: request.RoomId,
	}
	mobile, err := serverService.GetStatus(player, "one")
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	if mobile != tokenClaims.Mobile {
		mobile, err = serverService.GetStatus(player, "two")
		if err != nil {
			response := tool.Failure(resp, 500, "服务器错误")
			return response, err
		}
		if mobile != tokenClaims.Mobile {
			response := tool.Failure(resp, 400, "电话号码错误")
			return response, nil
		}
		player.PlayerId = "two"
		err = serverService.SetStatus(player, request)
		if err != nil {
			response := tool.Failure(resp, 500, "服务器错误")
			return response, err
		}
		resp = tool.Success(resp, 200, "状态设置成功")
		return resp, nil
	}
	player.PlayerId = "one"
	err = serverService.SetStatus(player, request)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	err = serverService.SetStatus(player, request)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	resp = tool.Success(resp, 200, "状态设置成功")
	return resp, nil
}

func (r *roomService) Exit(ctx context.Context, request *service.ExitRequest) (*service.Response, error) {
	resp := &service.Response{}

	return resp, nil
}
