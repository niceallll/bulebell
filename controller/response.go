package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 10001, //程序中的错误码
	"msg": xxx, 提示信息
	"data":{},
}
*/
type ResponseDate struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseDate{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}
func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseDate{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseDate{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
