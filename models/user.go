package models

import (
	"encoding/json"
	"errors"
)

type User struct {
	UserID       uint64 `json:"user_id,string" db:"user_id"` // 指定json序列化/反序列化时使用小写user_id
	Username     string `json:"username" db:"username"`
	Password     string `json:"password" db:"password"`
	Email        string `json:"email" db:"gender"`  // 邮箱
	Gender       int    `json:"gender" db:"gender"` // 性别
	AccessToken  string
	RefreshToken string
}
type VoteDataForm struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"`
}

func (v *VoteDataForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		Post_id   string `json:"post_id"`
		Direction int8   `json:"direction"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.Post_id) == 0 {
		err = errors.New("缺少必填字段post_id")
	} else if required.Direction == 0 {
		err = errors.New("缺少必填字段direction")
	} else {
		v.PostID = required.Post_id
		v.Direction = required.Direction
	}
	return
}
