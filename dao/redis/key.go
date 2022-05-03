package redis

const (
	KeyPostVotedZSetPf = "post:voted:" // zset;记录用户及投票的类型；参数是postid
	KeyPostTimeZSet    = "post:time"   //ZSet:帖子发帖时间
	KeyPostScoreZSet   = "post:score"  //ZSet:帖子投票分数
	KeyPrefix          = "bulebell:"
)

//给key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
