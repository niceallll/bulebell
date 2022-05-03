package logic

import (
	"bulebell/dao/mysql"
	"bulebell/models"
	"bulebell/pkg/jwt"
	"bulebell/pkg/snowflake"
)

// 存放业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生产UID，
	userId := snowflake.GenID()
	user := &models.User{
		UserID:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 密码加密
	// 保存进数据库
	return mysql.InsetUser(user)
}

func Login(p *models.LoginSignUp) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}
