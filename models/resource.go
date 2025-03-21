package models

import (
	"time"
)

// Resource 表示文件资源的模型
type Resource struct {
	ID         int64     `json:"id" db:"id"`
	Filename   string    `json:"filename" db:"filename"`
	Path       string    `json:"path" db:"path"`
	AuthorID   uint64    `json:"author_id" db:"author_id"`
	Status     int       `json:"status" db:"status"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
