package model

import (
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/db"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/redis"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/config"
)

// MessageTemplateKey 短信模板缓存key
const MessageTemplateKey = "message:template:merchant_id:%s:branch_id:%s"

var (
	entity      *db.Entity
	redisEntity *redis.Entity
)

// DatabaseConnection 数据库连接
func DatabaseConnection() (*db.Entity, error) {
	cfg := config.Setting.Database

	conn, err := db.NewEntity(db.Config{
		Dialect:  "postgres",
		Server:   cfg.Host,
		Port:     cfg.Port,
		User:     cfg.User,
		Database: cfg.DBName,
		Password: cfg.Password,
		Debug:    cfg.Log,
	})
	entity = conn
	return conn, err
}

// RedisConnection redis连接
func RedisConnection() (*redis.Entity, error) {
	conn, err := redis.NewEntity(redis.Config{
		IP:       config.Setting.Redis.IP,
		Port:     config.Setting.Redis.Port,
		Password: config.Setting.Redis.Password,
		Database: config.Setting.Redis.Database,
	})
	redisEntity = conn
	return conn, err
}
