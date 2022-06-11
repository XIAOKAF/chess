package api

import (
	"chess/server/server_service"
	"chess/service"
	"chess/tool"
	"context"
	"fmt"
)

var UserService = &userService{}

type userService struct {
}

func (u *userService) Register(ctx context.Context, request *service.RegisterRequest) (*service.Response, error) {
	resp := &service.Response{}
	//数据库查询电话号码是否存在,true表示存在
	err, flag := server_service.SelectUser(request.Mobile)
	fmt.Println(request)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	if flag {
		response := tool.Failure(resp, 400, "电话号码已被注册")
		return response, err
	}
	//获取验证码
	code, err := server_service.GetCode(request.Mobile)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	//校验验证码
	if request.Code != code {
		response := tool.Failure(resp, 400, "验证码错误或已过期")
		return response, nil
	}
	//注册
	err = server_service.InsertUser(request)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	response := tool.Success(resp, 200, "注册成功")
	return response, nil
}

func (u *userService) Login(ctx context.Context, request *service.LoginRequest) (*service.Response, error) {
	resp := &service.Response{}
	if request.Mobile == "" || request.Password == "" {
		response := tool.Failure(resp, 400, "必要字段不能为空")
		return response, nil
	}
	pwd, err := server_service.SelectPwd(request.Mobile)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	if pwd != request.Password {
		response := tool.Failure(resp, 400, "密码错误")
		return response, nil
	}
	err, token := server_service.CreateToken(request.Mobile, 2)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	err = server_service.StoreToken(request, token)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, nil
	}
	resp = tool.Failure(resp, 200, token)
	return resp, nil
}
