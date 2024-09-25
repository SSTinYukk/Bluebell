package mysql

import "errors"

var (
	ErrorInvalidID    = "无效的ID"
	ErrorInsertFailed = errors.New("插入数据失败")
)
