package redis

const (
	KeyPostInfoHashPrefix = "bluebell:post:"
	KeyPostTimeZSet       = "bluebell:post:time"  // zset;帖子及发帖时间定义
	KeyPostScoreZSet      = "bluebell:post:score" // zset;帖子及投票分数定义
	//KeyPostVotedUpSetPrefix   = "bluebell-plus:post:voted:down:"
	//KeyPostVotedDownSetPrefix = "bluebell-plus:post:voted:up:"
	KeyPostVotedZSetPrefix    = "bluebell:post:voted:" // zSet;记录用户及投票类型;参数是post_id
	KeyCommunityPostSetPrefix = "bluebell:community:"  // set保存每个分区下帖子的id
)
