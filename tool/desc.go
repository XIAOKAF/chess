package tool

import "github.com/gin-gonic/gin"

// Failure 返回错误
func Failure(info string, code int) {
	var ctx *gin.Context
	ctx.JSON(code, gin.H{
		"code": code,
		"info": info,
	})
}

// Success 返回正确
func Success(info string, code int) {
	var ctx *gin.Context
	ctx.JSON(code, gin.H{
		"code": code,
		"info": info,
	})
}
