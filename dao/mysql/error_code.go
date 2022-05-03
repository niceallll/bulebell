package mysql

import "errors"

var (
	ErrorUserExist    = errors.New("用户已存在")
	ErrorUserNotExist = errors.New("用户不存在")
	ErrorInvalidID    = errors.New("无效的id")
	ErrorPassword     = errors.New("密码错误")
)
