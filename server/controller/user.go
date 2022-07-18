package controller

import (
	"chess/proto"
	"chess/server/service"
	"chess/server/tool"
	"context"
)

var UserService = &userService{}

type userService struct {
}

func (u *userService) Register(ctx context.Context, request *proto.RegisterRequest) (*proto.Response, error) {
	resp := &proto.Response{}
	//数据库查询电话号码是否存在,true表示存在
	err, flag := service.SelectUser(request.Mobile)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	if flag {
		response := tool.Failure(resp, 400, "电话号码已被注册")
		return response, err
	}
	//获取验证码
	code, err := service.GetCode(request.Mobile)
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
	err = service.InsertUser(request)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	response := tool.Success(resp, 200, "注册成功")
	return response, nil
}

func (u *userService) Login(ctx context.Context, request *proto.LoginRequest) (*proto.Response, error) {
	resp := &proto.Response{}
	if request.Mobile == "" || request.Password == "" {
		response := tool.Failure(resp, 400, "必要字段不能为空")
		return response, nil
	}
	pwd, err := service.SelectPwd(request.Mobile)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	if pwd != request.Password {
		response := tool.Failure(resp, 400, "密码错误")
		return response, nil
	}
	err, token := service.CreateToken(request.Mobile, 2)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, err
	}
	err = service.StoreToken(request, token)
	if err != nil {
		response := tool.Failure(resp, 500, "服务器错误")
		return response, nil
	}
	resp = tool.Failure(resp, 200, token)
	return resp, nil
}
