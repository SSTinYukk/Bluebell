package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// UploadResource 处理文件上传逻辑
// UploadResource 处理文件上传逻辑
func UploadResource(fileData []byte, filename string, authorID uint64) (*models.Resource, error) {
	// 生成唯一的文件名
	uniqueID, _ := snowflake.GenID()
	ext := filepath.Ext(filename)
	newFilename := fmt.Sprintf("%d%s", uniqueID, ext)

	// 定义文件存储路径
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			return nil, err
		}
	}
	filePath := filepath.Join(uploadDir, newFilename)

	// 保存文件到本地
	if err := os.WriteFile(filePath, fileData, 0644); err != nil {
		return nil, err
	}

	now := time.Now()
	// 创建资源对象
	resource := &models.Resource{
		Filename:   newFilename,
		Path:       filePath,
		AuthorID:   authorID,
		Status:     0, // 初始状态为待审核
		CreateTime: now,
		UpdateTime: now,
	}

	// 插入资源信息到数据库
	if err := mysql.InsertResource(resource); err != nil {
		return nil, err
	}

	return resource, nil
}

// GetResourceByID 根据资源 ID 获取文件资源信息
func GetResourceByID(id int64) (*models.Resource, error) {
	return mysql.GetResourceByID(id)
}
