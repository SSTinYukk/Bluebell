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

func GetPostByID(pid uint64) (*models.Post, error) {
	return mysql.GetPostByID(pid)
}
