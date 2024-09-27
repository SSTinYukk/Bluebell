package mysql

import (
	"bluebell/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id, title, content, author_id, community_id)
		values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.PostID, p.Title, p.Content, p.AuthorId, p.CommunityID)
	if err != nil {
		err = ErrorInsertFailed
		return
	}
	return
}

func GetPostByID(pid uint64) (p *models.Post, err error) {
	p = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time 
	from post 
	where post_id= ? `
	err = db.Get(p, sqlStr, pid)
	if err != nil {
		return
	}
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
		from post
		ORDER BY create_time
		DESC
		limit ?,?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
