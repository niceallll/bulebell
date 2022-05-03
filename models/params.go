package models

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// LoginSignUp登录参数
type LoginSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//投票数据
type VoteData struct {
	//从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`              //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成（1）还是反对（-1）取消（0）
}
