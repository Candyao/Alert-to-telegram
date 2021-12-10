package util

import (
	"github.com/gin-gonic/gin"
)

type GinSelf struct {
	Ctx *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *GinSelf) Response(httpCode, errCode int, data interface{}) {
	g.Ctx.JSON(httpCode, Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
	return
}
