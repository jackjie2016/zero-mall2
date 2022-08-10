package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-mal/global"
	model "zero-mal/service/user/model/genModel"
	GormModel "zero-mal/service/user/model/gorm"
	"zero-mal/service/user/rpc/initialize"
	"zero-mal/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel     model.UserModel
	UserGormModel GormModel.UserModel
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
	}
}
