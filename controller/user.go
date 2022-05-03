package controller

import (
	"bulebell/dao/mysql"
	"bulebell/logic"
	"bulebell/models"
	"bulebell/settings"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1，获取参数和参数校验
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2，业务处理
	fmt.Println(p)
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic Signup failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3，返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	//1,获取请求参数及参数校验
	p := new(models.LoginSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2，业务逻辑处理

	token, err := logic.Login(p)
	if err != nil {

		zap.L().Error("Login.login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
		//3，返回响应
	}
	ResponseSuccess(c, token)
}

func Version(c *gin.Context) {
	s := settings.Conf.Version
	fmt.Println(s)
}
