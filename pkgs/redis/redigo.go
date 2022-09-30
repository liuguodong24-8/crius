package redis

import (
	"errors"
	"fmt"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/util"

	redigo "github.com/gomodule/redigo/redis"
)

// DEL del
func (e *Entity) DEL(key ...interface{}) error {
	conn, err := e.getRedisConn()
	if err != nil {
		return err
	}

	_, err = conn.Do("DEL", key...)
	return err
}

// HSET hset
func (e *Entity) HSET(key, field, value interface{}) error {
	conn, err := e.getRedisConn()
	if err != nil {
		return err
	}

	_, err = conn.Do("HSET", key, field, value)
	return err
}

// HGET hget
func (e *Entity) HGET(key, field interface{}) ([]byte, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return nil, err
	}

	reply, err := conn.Do("HGET", key, field)

	return redigo.Bytes(reply, err)
}

// HEXISTS hexists
func (e *Entity) HEXISTS(key, field interface{}) (bool, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return false, err
	}
	reply, err := conn.Do("HEXISTS", key, field)

	return redigo.Bool(reply, err)
}

// HDEL hdel
func (e *Entity) HDEL(key, field interface{}) error {
	conn, err := e.getRedisConn()
	if err != nil {
		return err
	}
	_, err = conn.Do("HDEL", key, field)

	return err
}

// HGETALL hgetall
func (e *Entity) HGETALL(key interface{}) ([][]byte, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return nil, err
	}
	reply, err := conn.Do("HGETALL", key)

	return redigo.ByteSlices(reply, err)
}

// HMSET hmset, args type is map or struct
func (e *Entity) HMSET(key interface{}, args interface{}) error {
	conn, err := e.getRedisConn()
	if err != nil {
		return err
	}

	_, err = conn.Do("HMSET", redigo.Args{}.Add(key).AddFlat(args)...)

	return err
}

// HGETALLSTRUCT hgetall, but return dest as a struct
func (e *Entity) HGETALLSTRUCT(key interface{}, dest interface{}) error {
	conn, err := e.getRedisConn()
	if err != nil {
		return err
	}
	values, err := redigo.Values(conn.Do("HGETALL", key))
	if err != nil {
		return err
	}
	return redigo.ScanStruct(values, dest)
}

// HVALS hvals
func (e *Entity) HVALS(key interface{}) ([][]byte, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return nil, err
	}
	reply, err := conn.Do("HVALS", key)

	return redigo.ByteSlices(reply, err)
}

// HMGET hmget
func (e *Entity) HMGET(key interface{}, args []interface{}) ([][]byte, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return nil, err
	}

	var command []interface{}
	command = append(command, key)
	for _, v := range args {
		command = append(command, v)
	}

	reply, err := conn.Do("HMGET", command...)

	return redigo.ByteSlices(reply, err)
}

// INCR ...
func (e *Entity) INCR(key interface{}) (int64, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return 0, err
	}

	reply, err := conn.Do("INCR", key)
	return redigo.Int64(reply, err)
}

// EXPIRE ...
func (e *Entity) EXPIRE(key, ttl interface{}) error {
	conn, err := e.getRedisConn()
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, ttl)
	return err
}

// GET ...
func (e *Entity) GET(key interface{}) (string, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return "", err
	}

	reply, err := conn.Do("GET", key)
	return redigo.String(reply, err)
}

// SET ...
func (e *Entity) SET(key, value interface{}) error {
	conn, err := e.getRedisConn()
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	return err
}

// ZADD zadd
func (e *Entity) ZADD(key interface{}, scoreValues map[int64]interface{}) error {
	conn, err := e.getRedisConn()
	if err != nil {
		return err
	}

	var command []interface{}
	command = append(command, key)
	for score, value := range scoreValues {
		command = append(command, score, value)
	}

	_, err = conn.Do("ZADD", command...)
	return err
}

// ZREM zrem
func (e *Entity) ZREM(key interface{}, members ...interface{}) (int64, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return 0, err
	}

	var command []interface{}
	command = append(command, key)
	command = append(command, members...)
	return redigo.Int64(conn.Do("ZREM", command...))
}

// ZCARD zcard
func (e *Entity) ZCARD(key interface{}) (int64, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return 0, err
	}

	return redigo.Int64(conn.Do("ZCARD", key))
}

// ZSCORE zscore
func (e *Entity) ZSCORE(key, member interface{}) (float64, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return 0, err
	}

	return redigo.Float64(conn.Do("ZSCORE", key, member))
}

// ZRANGEWITHSCORES zrange
func (e *Entity) ZRANGEWITHSCORES(key interface{}, min, max int64) (map[string]string, error) {
	conn, err := e.getRedisConn()
	if err != nil {
		return nil, err
	}

	return redigo.StringMap(conn.Do("ZRANGE", key, min, max, "withscores"))
}

func (e *Entity) getRedisConn() (redigo.Conn, error) {
	var conn redigo.Conn
	var err error
	conn, err = e.redis.Dial()
	if err == nil {
		return conn, nil
	}

	util.Logger.Error(fmt.Sprintf("捕获redis错误信息:%s", err.Error()))
	var maxRetry = 3
	for i := 0; i < maxRetry; i++ {
		pool := newPool(e.config)
		e.redis = pool
		conn := pool.Get()
		_, err := conn.Do("PING")
		if err != nil {
			return conn, nil
		}
	}

	return nil, errors.New("redis error more than 3 times")
}

func newPool(cfg Config) *redigo.Pool {
	addr := fmt.Sprintf("%s:%d", cfg.IP, cfg.Port)

	return &redigo.Pool{
		MaxIdle:     100,
		MaxActive:   4000,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", addr, redigo.DialPassword(cfg.Password), redigo.DialDatabase(cfg.Database))
			if nil != err {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
