package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"zero-mal/service/common/api/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Cache cache.Cache //缓存

}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cache:  cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("dc"), nil),
	}
}
