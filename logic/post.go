package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	postID, err := snowflake.GenID()
	if err != nil {
		return err
	}
	p.PostID = postID
	if err := mysql.CreatePost(p); err != nil {
		zap.L().Error("mysql.CreatePost(p) failed", zap.Error(err))
		return err
	}
	return nil
}

func GetPostByID(pid uint64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID failed",
			zap.Uint64("postID", pid),
			zap.Error(err))
		return
	}
	//根据作者id差查询作者信息
	user, err := mysql.GetUserByID(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(pid) failed",
			zap.Uint("postid", uint(post.AuthorId)),
			zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed",
			zap.Uint64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		Post:               post,
		CommunityDetailRes: community,
		AuthorName:         user.Username,
	}
	return
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList failed")
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		user, err := mysql.GetUserByID(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(pid) failed",
				zap.Uint("postid", uint(post.AuthorId)),
				zap.Error(err))
			continue
		}
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityByID() failed",
				zap.Uint64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		tmp := &models.ApiPostDetail{
			Post:               post,
			CommunityDetailRes: community,
			AuthorName:         user.Username,
		}
		data = append(data, tmp)
	}
	return
}
