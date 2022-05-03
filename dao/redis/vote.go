package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

/* 投票的情况：
direction = 1; 两种情况
	1，之前没有投票，现在投赞成票        --- > 更新分数和投票记录  差值的绝对值： 1 + 432
	2，之前投反对票，现在投赞成票        --- > 更新分数和投票记录  差值的绝对值： 2 + 432*2
direction = 0；两种情况
    1，之前投反对票，现在要取消投票      --- > 更新分数和投票记录  差值的绝对值： 1 + 432
	2，之前投赞成票，现在要取消投票      --- > 更新分数和投票记录  差值的绝对值： 1 - 432
dierction = -1；两种情况
	1，之前没有投票，现在投反对票        --- > 更新分数和投票记录  差值的绝对值： 1 -432
	2，之前投赞成票，现在投反对票        --- > 更新分数和投票记录  差值的绝对值： 2 -432*2
投票的限制：
每个帖子自发表之日起一个星期内允许用户投票，超过一个星期就不允许再投票。
	1，到期之后将redis中保存的赞成票和反对票保存到mysql
	2，到期之后删除 KeyPostVotedZSetPf
*/
const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票的分数
)

var (
	ErrVoteTimeExpore = errors.New("投票时间已过")
)

func VoteForPost(userId, postID string, value float64) error {
	// 判断投票限制
	// 去redis取出时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpore
	}
	// 更新帖子的分数
	// 查询当前用户给当前帖子的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZSetPf+postID), userId).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算两次的差值
	_, err := rdb.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID).Result()
	if ErrVoteTimeExpore != nil {
		return err
	}

	// 记录用户为帖子投过票
	if value == 0 {
		rdb.ZRem(getRedisKey(KeyPostVotedZSetPf+postID), userId).Result()
	} else {
		_, err = rdb.ZAdd(getRedisKey(KeyPostVotedZSetPf+postID), redis.Z{
			value,
			userId,
		}).Result()
	}
	return err
}
