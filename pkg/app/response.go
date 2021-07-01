package app

import (
	"admin/pkg/code"
	"github.com/gin-gonic/gin"
)

type Response struct {
	C *gin.Context
}

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Response) JSONResponse(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Res{
		Code: errCode,
		Msg:  code.GetMsg(errCode),
		Data: data,
	})
	return
}
