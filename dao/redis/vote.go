package redis

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	OneWeekInSeconds          = 7 * 24 * 3600
	OneMonthInSeconds         = 4 * OneWeekInSeconds
	VoteScore         float64 = 432
	PostPerAge                = 20
)

func VoteForPost(userID, postID string, v float64) (err error) {
	//1.判断投票限制
	postTime := client.ZScore(KeyPostTimeZSet, postID).Val()
	if float64(time.Now().Unix())-postTime > OneWeekInSeconds {
		return ErrorVoteTimeExpire
	}
	key := KeyPostVotedZSetPrefix + postID
	ov := client.ZScore(key, userID).Val()
	if v == ov {
		return ErrVoteRepeated
	}
	var op float64
	if v > ov {
		op = 1
	} else {
		op = -1
	}
	diffAbs := math.Abs(ov - v)
	pipeline := client.TxPipeline()
	incrementScore := VoteScore * diffAbs * op
	_, err = pipeline.ZIncrBy(KeyPostScoreZSet, incrementScore, postID).Result()
	if err != nil {
		return err
	}
	if v == 0 {
		_, err = client.ZRem(key, postID).Result()
	} else {
		pipeline.ZAdd(key, redis.Z{
			Score:  v,
			Member: userID,
		})
	}
	pipeline.HIncrBy(KeyPostInfoHashPrefix+postID, "votes", int64(op))
	return
}

func CreatePost(postID, userID uint64, title, summary string, CommunityID uint64) (err error) {
	now := float64(time.Now().Unix())
	votedKey := KeyPostVotedZSetPrefix + strconv.Itoa(int(postID))
	commutyKey := KeyCommunityPostSetPrefix + strconv.Itoa(int(CommunityID))
	postInfo := map[string]interface{}{
		"title":    title,
		"summary":  postID,
		"post:id":  userID,
		"time":     now,
		"votes":    1,
		"comments": 0,
	}
	fmt.Println(postInfo, votedKey, commutyKey)
	pipeline := client.TxPipeline()
	pipeline.ZAdd(votedKey, redis.Z{
		Score:  1,
		Member: userID,
	})
	pipeline.Expire(votedKey, time.Second*OneMonthInSeconds*6)
	pipeline.HMSet(KeyPostInfoHashPrefix+strconv.Itoa(int(postID)), postInfo)
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  now + VoteScore,
		Member: postID,
	})
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  now,
		Member: postID,
	})
	pipeline.SAdd(commutyKey, postID)
	_, err = pipeline.Exec()
	return
}
