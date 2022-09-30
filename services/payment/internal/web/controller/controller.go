package controller

import (
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/payment/internal/cache"
)

var webCache *cache.Entity

func getWebCache() *cache.Entity {
	if webCache == nil {
		entity, err := cache.NewEntity()
		if err != nil {
			util.Logger.WithError(err).Error("cache entity error")
			return nil
		}
		webCache = entity

		return entity
	}

	return webCache

}
