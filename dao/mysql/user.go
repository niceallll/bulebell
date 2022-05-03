package mysql

import (
	"bulebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

// 把每一步数据库操作封装成函数
// 待logic层更具业务调用

const secret = "wenchao.ma"

// CheckUserExist 检查用户名是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsetUser 向数据库中插入一条新的用户记录
func InsetUser(user *models.User) (err error) {
	//对密码加密
	user.Password = encyptPassword(user.Password)
	// 执行sql语句入库
	sqlStr := `insert into user(user_id,username,password)values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// encyptPassword 对用户密码加密
func encyptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	//h.Sum([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password //用户登录的密码
	sqlStr := `select user_id, username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	// 判断密码是否正确
	password := encyptPassword(oPassword)
	if password != user.Password {
		return ErrorPassword
	}
	return
}
