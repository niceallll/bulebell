package controller

import (
	"bulebell/logic"
	"bulebell/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteC(c *gin.Context) {
	//参数校验
	p := new(models.VoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("post vote with invalid param")
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("VoteForPost", zap.Error(err))
	}
	ResponseSuccess(c, nil)
}
