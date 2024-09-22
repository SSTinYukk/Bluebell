package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func GetCommunityByID(id uint64) (*models.CommunityDetailRes, error) {
	commty := new(models.CommunityDetail)
	sqlStr := "select community_id,community_name,introduction,create_time from community where community_id = ?"
	err := db.Get(commty, sqlStr, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(ErrorInvalidID)
		}
		zap.L().Error("query community failed", zap.String("sql", sqlStr), zap.Error(err))
	}
	return &models.CommunityDetailRes{
		CommunityID:   commty.CommunityID,
		CommunityName: commty.CommunityName,
		Introduction:  commty.Introduction,
		CreateTime:    commty.CreateTime.Format("2006-01-02 15:04:05"),
	}, err
}
