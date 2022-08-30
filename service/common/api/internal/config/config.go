package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	AliSmsInfo struct {
		ApiKey       string
		ApiSecret    string
		SignName     string
		TemplateCode string
	}
	Cache      cache.CacheConf
	CacheRedis cache.CacheConf
}
