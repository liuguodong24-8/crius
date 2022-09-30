package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// LinkDetail 详情
type LinkDetail struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// Link 配置跳转链接信息
type Link struct {
	Official    []LinkDetail `json:"official"`
	Miniprogram []LinkDetail `json:"miniprogram"`
}

// CacheValue 缓存value
func (l Link) CacheValue() ([]byte, error) {
	return json.Marshal(l)
}

// CacheScan 缓存scan
func (l *Link) CacheScan(src []byte) error {
	return json.Unmarshal(src, l)
}

// GetWechatLink 获取微信链接配置
func GetWechatLink(ctx context.Context, cache *Entity) (Link, error) {
	cacheKey := fmt.Sprintf("wechat_link")
	item, cacheErr := cache.Get(cacheKey)
	var link Link
	if nil == cacheErr {
		if err := link.CacheScan(item); nil == err {
			return link, nil
		}
	}

	res, linkErr := getLinkConfigFile()
	if linkErr != nil {
		return link, linkErr
	}

	cacheVal, valErr := res.CacheValue()
	if valErr != nil {
		return link, valErr
	}

	if err := cache.Set(cacheKey, cacheVal); nil != err {
		util.Logger.WithMetadata(ctx).WithError(err).Error("设置内存缓存失败")
	}

	return res, nil
}

func getLinkConfigFile() (Link, error) {
	var items Link

	path, err := os.Getwd()
	if err != nil {
		return items, errors.New("获取配置文件错误")
	}

	file := fmt.Sprintf("%s/config/link.json", path)
	exists, err := util.FileExists(file)
	if err != nil {
		return items, errors.New("获取配置文件错误, 校验文件是否存在")
	}
	if !exists {
		return items, errors.New("获取配置文件错误, 不存在")
	}

	bytes, readErr := ioutil.ReadFile(file)
	if readErr != nil {
		return items, fmt.Errorf("读取配置文件错误:%s", readErr.Error())
	}

	if err := json.Unmarshal(bytes, &items); nil != err {
		return items, fmt.Errorf("映射配置文件错误:%s", err.Error())
	}

	return items, nil
}
