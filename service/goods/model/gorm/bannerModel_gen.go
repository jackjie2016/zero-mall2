// Code generated by goctl. DO NOT EDIT!

package gorm

import (
	"context"
	"fmt"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stringx"
	"gorm.io/gorm"
	"strings"
)

var (
	bannerFieldNames          = builder.RawFieldNames(&Banner{})
	bannerRows                = strings.Join(bannerFieldNames, ",")
	bannerRowsExpectAutoSet   = strings.Join(stringx.Remove(bannerFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bannerRowsWithPlaceHolder = strings.Join(stringx.Remove(bannerFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheGoZeroMallBannerIdPrefix = "cache:goZeroMall:banner:id:"
)

type (
	bannerModel interface {
		Insert(ctx context.Context, data *Banner) error
		FindOne(ctx context.Context, id int64) (*Banner, error)
		Update(ctx context.Context, data *Banner) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBannerModel struct {
		gormc.CachedConn
		table string
	}

	Banner struct {
		BaseModel

		Image string `gorm:"type:varchar(200);not null"`
		Url   string `gorm:"type:varchar(200);not null"`
		Index int32  `gorm:"type:int;default:1;not null"`
	}
)

func newBannerModel(conn *gorm.DB, c cache.CacheConf) *defaultBannerModel {
	return &defaultBannerModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`banner`",
	}
}

func (m *defaultBannerModel) Insert(ctx context.Context, data *Banner) error {
	goZeroMallBannerIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallBannerIdPrefix, data.Id)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Save(&data).Error
	}, goZeroMallBannerIdKey)
	return err
}

func (m *defaultBannerModel) FindOne(ctx context.Context, id int64) (*Banner, error) {
	goZeroMallBannerIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallBannerIdPrefix, id)
	var resp Banner
	err := m.QueryCtx(ctx, &resp, goZeroMallBannerIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Banner{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBannerModel) Update(ctx context.Context, data *Banner) error {
	goZeroMallBannerIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallBannerIdPrefix, data.Id)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Save(data).Error
	}, goZeroMallBannerIdKey)
	return err
}

func (m *defaultBannerModel) Delete(ctx context.Context, id int64) error {
	goZeroMallBannerIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallBannerIdPrefix, id)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Delete(&Banner{}, id).Error
	}, goZeroMallBannerIdKey)
	return err
}

func (m *defaultBannerModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGoZeroMallBannerIdPrefix, primary)
}

func (m *defaultBannerModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&Banner{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultBannerModel) tableName() string {
	return m.table
}
