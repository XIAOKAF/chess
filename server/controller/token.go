package controller

import (
	"chess-room/server/serverService"
	"chess-room/service"
	"chess-room/tool"
	"context"
)

func (r *roomService) JudgeToken(ctx context.Context, request *service.BasicRequest) (*service.Response, error) {
	resp := &service.Response{}
	if request.Token == "" {
		response := tool.Failure(resp, 400, "请先登录")
		return response, nil
	}
	tokenClaims, err := serverService.ParseToken(request.Token)
	if err != nil {
		if err.Error() == "fail to parse token" {
			response := tool.Failure(resp, 500, "服务器错误")
			return response, err
		}
		response := tool.Failure(resp, 400, "请先登录")
		return response, nil
	}
	tokenString, err := serverService.GetToken(tokenClaims)
	if tokenString != request.Token {
		response := tool.Failure(resp, 400, "token错误")
		return response, nil
	}
	return resp, nil
}
