package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityByID(id uint64) (*models.CommunityDetailRes, error) {
	return mysql.GetCommunityByID(id)
}
