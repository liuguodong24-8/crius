package cache

import (
	"errors"
	"time"

	"github.com/allegro/bigcache/v3"
)

var _ Interface = (*Entity)(nil)

// Interface 缓存接口定义
type Interface interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}

// NotFoundError ...
var NotFoundError error

// Entity 缓存实例
type Entity struct {
	*bigcache.BigCache
}

// NewEntity 实例化缓存
func NewEntity() (*Entity, error) {
	c, err := bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))

	return &Entity{c}, err
}

// Set 设置
func (e *Entity) Set(key string, value []byte) error {
	return e.BigCache.Set(key, value)
}

// Get 获取
func (e *Entity) Get(key string) ([]byte, error) {
	data, err := e.BigCache.Get(key)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, NotFoundError
	}

	return data, err
}
