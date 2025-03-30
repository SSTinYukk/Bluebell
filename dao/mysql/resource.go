package mysql

import (
	"bluebell/models"
)

// InsertResource 插入文件资源信息到数据库
func InsertResource(resource *models.Resource) error {
	sqlStr := `INSERT INTO resources (filename, path, author_id, status) VALUES (?,?,?,?)`
	result, err := db.Exec(sqlStr, resource.Filename, resource.Path, resource.AuthorID, resource.Status)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	resource.ID = id
	return nil
}

// GetResourceByID 根据资源 ID 获取文件资源信息
func GetResourceByID(id int64) (*models.Resource, error) {
	var resource models.Resource
	sqlStr := `SELECT id, filename, path, author_id, status, create_time, update_time 
               FROM resources WHERE id =?`
	err := db.Get(&resource, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func UpdateResourceStatus(id int64, status int) error {
	sqlStr := `UPDATE resources SET status =?, update_time = NOW() WHERE id =?`
	_, err := db.Exec(sqlStr, status, id)
	return err
}

func DeleteResource(id int64) error {
	sqlStr := `DELETE FROM resources WHERE id =?`
	_, err := db.Exec(sqlStr, id)
	return err
}
