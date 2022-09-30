package service

import (
	"fmt"

	"github.com/fideism/golang-wechat/cache"
	"github.com/fideism/golang-wechat/officialaccount"
	officialConfig "github.com/fideism/golang-wechat/officialaccount/config"
	"gitlab.omytech.com.cn/micro-service/wechat/internal/config"
)

// ErrcodeSuccess 微信成功码
const ErrcodeSuccess int64 = 0

// OfficialConfig 公众号配置
type OfficialConfig struct {
	AppID     string
	AppSecret string
	Token     string
	AesKey    string
}

// NewOfficial 构建公会号对象
func NewOfficial(channel string) (*officialaccount.OfficialAccount, OfficialConfig, error) {
	officialCfg, err := GetOfficialConfig(channel)
	if err != nil {
		return nil, officialCfg, err
	}

	redisCfg := config.Setting.Redis
	//设置全局cache，也可以单独为每个操作实例设置
	rds := &cache.RedisOpts{
		Host:     fmt.Sprintf("%s:%d", redisCfg.IP, redisCfg.Port),
		Database: redisCfg.Database,
		Password: redisCfg.Password,
	}

	return officialaccount.NewOfficialAccount(&officialConfig.Config{
		AppID:          officialCfg.AppID,
		AppSecret:      officialCfg.AppSecret,
		Token:          officialCfg.Token,
		EncodingAESKey: officialCfg.AesKey,
		Cache:          cache.NewRedis(rds),
	}), officialCfg, nil
}

// GetOfficialConfig 获取公众号配置
func GetOfficialConfig(channel string) (OfficialConfig, error) {
	for _, v := range config.Setting.Wechat.Official {
		if v.Channel == channel {
			return OfficialConfig{
				AppID:     v.AppID,
				AppSecret: v.AppSecret,
				Token:     v.Token,
				AesKey:    v.AesKey,
			}, nil
		}
	}

	return OfficialConfig{}, fmt.Errorf("%s对应公众号配置不存在", channel)
}
