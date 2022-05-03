package logic

import (
	"bulebell/dao/redis"
	"bulebell/models"
	"go.uber.org/zap"
	"strconv"
)

// 投票功能：
// 1，用户投票数据

func VoteForPost(userID int64, p *models.VoteData) error {
	zap.L().Debug("VoteForPost faild",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
	// 判断投票限制
	// 更新帖子的分数
	// 记录用户为帖子投过票
}
