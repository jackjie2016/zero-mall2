package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache   cache.CacheConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Rocketmq struct {
		Hosts []string
	}

	//UserRpc      zrpc.RpcClientConf
	InventoryRpc zrpc.RpcClientConf
	GoodsRpc     zrpc.RpcClientConf
}