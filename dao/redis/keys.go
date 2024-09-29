package redis

const (
	KeyPostInfoHashPrefix     = "bluebell"
	KeyPostTimeZSet           = "bluebell:post:time"        //zset;帖子的发帖时间定义
	KeyPostScore              = "bluebell:post:score"       //zset;帖子及投票分数定义
	KeyPostVotedZSetPrefix    = "bluebell-plus:post:voted:" //zset;记录用户及投票类型；参数是post_id
	KeyCommunityPostSetPredix = "bluebell-plus:community:"  //zset保存每个分区下帖子的id
)
