package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
	"zero-mal/global"
	model "zero-mal/service/user/model/genModel"
	GormModel "zero-mal/service/user/model/gorm"
	"zero-mal/service/user/rpc/internal/config"
	"zero-mal/service/user/rpc/internal/initialize"
)

type ServiceContext struct {
	Config config.Config

	UserModel     model.UserModel
	UserGormModel GormModel.UserModel

	Cache cache.Cache //缓存
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	if global.DB == nil {
		initialize.InitDb(c.DB.DataSource)
	}

	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlConn, c.Cache),
		UserGormModel: GormModel.NewUserModel(global.DB, c.Cache),
		Cache:         cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("dc"), nil),
	}
}
