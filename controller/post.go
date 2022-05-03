package controller

import (
	"bulebell/logic"
	"bulebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//
func CreatePostHandler(c *gin.Context) {
	//1, 获取参数以及参数的校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c中取到发出请求的id
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	p.AuthorID = userID
	//2，创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//3，返回响应
	ResponseSuccess(c, nil)
}

func GetPostHandler(c *gin.Context) {
	//获取参数，（从url中帖子的id）
	pidstr := c.Param("id")
	pid, err := strconv.ParseInt(pidstr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic GetPostById(pid) fiald", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	} //根据id取出帖子数据
	//返回响应
	ResponseSuccess(c, data)
}
