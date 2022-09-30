package redis

import (
	redigo "github.com/gomodule/redigo/redis"
)

// 强制要求该接口方法必须实现
var _ Interface = (*Entity)(nil)

// Interface redis接口定义
type Interface interface {
	HSET(key, field, value interface{}) error
	HGET(key, field interface{}) ([]byte, error)
	HEXISTS(key, field interface{}) (bool, error)
	HDEL(key, field interface{}) error
	HGETALL(key interface{}) ([][]byte, error)
	HVALS(key interface{}) ([][]byte, error)
	HMGET(key interface{}, keys []interface{}) ([][]byte, error)
	INCR(key interface{}) (int64, error)
	EXPIRE(key, ttl interface{}) error
	SET(key, value interface{}) error
	GET(key interface{}) (string, error)
}

// Config redis config
type Config struct {
	IP       string
	Port     int
	Password string
	Database int
}

// Entity redis 实例
type Entity struct {
	redis  *redigo.Pool
	config Config
}

// NewEntity 实例化 redis
func NewEntity(cfg Config) (*Entity, error) {
	return &Entity{
		redis:  newPool(cfg),
		config: cfg,
	}, nil
}
