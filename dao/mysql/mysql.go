package mysql

import (
	"bluebell/settings"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.AppConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=true&loc=Local", cfg.MySQLConfig.User, cfg.MySQLConfig.Password, cfg.MySQLConfig.Host, cfg.MySQLConfig.Port, cfg.MySQLConfig.DB)
	if db, err = sqlx.Connect("mysql", dsn); err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}
