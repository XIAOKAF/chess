package tool

import (
	"chess/service"
)

// Failure 返回错误
func Failure(response *service.Response, code int32, info string) *service.Response {
	response.Code = code
	response.Info = info
	return response
}

// Success 返回正确
func Success(response *service.Response, code int32, info string) *service.Response {
	response.Code = code
	response.Info = info
	return response
}
