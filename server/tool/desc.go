package tool

import (
	"chess/proto"
)

// Failure 返回错误
func Failure(response *proto.Response, code int32, info string) *proto.Response {
	response.Code = code
	response.Info = info
	return response
}

// Success 返回正确
func Success(response *proto.Response, code int32, info string) *proto.Response {
	response.Code = code
	response.Info = info
	return response
}
