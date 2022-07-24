package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"zero-mal/service/user/api/internal/config"
	"zero-mal/service/user/api/internal/middleware"
	"zero-mal/service/user/rpc/usercenter"
)

type ServiceContext struct {
	Config  config.Config
	IsLogin rest.Middleware
	IsAdmin rest.Middleware
	UserRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		IsLogin: middleware.NewIsLoginMiddleware().Handle,
		IsAdmin: middleware.NewIsAdminMiddleware().Handle,
		UserRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpc)),
	}
}
