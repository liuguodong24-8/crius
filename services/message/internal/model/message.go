package model

import (
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/db"
	"gitlab.omytech.com.cn/micro-service/message/internal/config"
)

// DatabaseConnection 数据库连接
func DatabaseConnection() (*db.Entity, error) {
	cfg := config.Setting.Database

	return db.NewEntity(db.Config{
		Dialect:  "postgres",
		Server:   cfg.Host,
		Port:     cfg.Port,
		User:     cfg.User,
		Database: cfg.DBName,
		Password: cfg.Password,
		Debug:    cfg.Log,
	})
}
