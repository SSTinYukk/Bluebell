package mysql

import (
	"bluebell/models"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id, title, content, author_id, community_id)
		values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorId, p.CommunityID)
	if err != nil {
		zap.L().Error("insert post failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}
func GetPostByID(pid uint64) (p *models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time 
	from post 
	where post_id= ? `
	db.Get(post, sqlStr, pid)
}
