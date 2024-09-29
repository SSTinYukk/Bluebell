package controller

import (
	"bluebell/dao/redis"
	"bluebell/logic"
	"bluebell/models"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type VoteData struct {
	PostID    string `json:"post_id,string"`
	Direction int    `json:"direction,string"`
}

func (v *VoteData) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		PostID    string `json:"post_id"`
		Direction int    `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return

	} else if len(required.PostID) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.PostID
		v.Direction = required.Direction
	}
	return
}

func VoteHandler(c *gin.Context) {
	vote := new(models.VoteDataForm)
	if err := c.ShouldBindJSON(&vote); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParams, errData)
		return
	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
	}
	if err := logic.VoteForPost(userID, vote); err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		switch err {
		case redis.ErrVoteRepeated:
			ResponseError(c, ErrVoteRepeated)
		case redis.ErrorVoteTimeExpire:
			ResponseError(c, ErrorVoteTimeExpire)
		default:
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, nil)
}
