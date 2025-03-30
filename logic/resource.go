package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"errors"
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

func GetResourceByID(id int64) (*models.Resource, error) {
	resource, err := mysql.GetResourceByID(id)
	if err != nil {
		return nil, err
	}
	// 检查资源状态，只有状态为 1（已通过）才允许访问
	if resource.Status != 1 {
		return nil, errors.New("资源未通过审核，无法访问")
	}
	return resource, nil
}

func ReviewResource(id int64, status int) error {
	// 简单检查状态值是否合法，假设 0 为待审核，1 为通过，2 为拒绝
	if status < 0 || status > 2 {
		return errors.New("无效的审核状态")
	}
	return mysql.UpdateResourceStatus(id, status)
}

func DeleteResource(id int64) error {
	// 获取资源信息
	resource, err := mysql.GetResourceByID(id)
	if err != nil {
		return err
	}

	// 删除本地文件
	err = os.Remove(resource.Path)
	if err != nil {
		return err
	}

	// 删除数据库记录
	return mysql.DeleteResource(id)
}
