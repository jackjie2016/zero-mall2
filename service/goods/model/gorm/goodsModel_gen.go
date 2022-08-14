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
	goodsFieldNames          = builder.RawFieldNames(&Goods{})
	goodsRows                = strings.Join(goodsFieldNames, ",")
	goodsRowsExpectAutoSet   = strings.Join(stringx.Remove(goodsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	goodsRowsWithPlaceHolder = strings.Join(stringx.Remove(goodsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheGoZeroMallGoodsIdPrefix = "cache:goZeroMall:goods:id:"
)

type (
	goodsModel interface {
		Insert(ctx context.Context, data *Goods) error
		FindOne(ctx context.Context, id int64) (*Goods, error)
		Update(ctx context.Context, data *Goods) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGoodsModel struct {
		gormc.CachedConn
		table string
	}

	Goods struct {
		BaseModel

		CategoryID int32 `gorm:"type:int;not null;column:category_id"`
		Category   Category
		BrandsID   int32  `gorm:"type:int;not null;column:brand_id"`
		Brands     Brands `gorm:"foreignKey:BrandsID;references:Id" json:"-"`

		OnSale          bool     `gorm:"column:on_sale"`
		GoodsSn         string   `gorm:"column:goods_sn"`
		Name            string   `gorm:"column:name"`
		ClickNum        int32    `gorm:"column:click_num"`
		SoldNum         int32    `gorm:"column:sold_num"`
		FavNum          int32    `gorm:"column:fav_num"`
		Stocks          int32    `gorm:"column:stocks"`
		MarketPrice     float32  `gorm:"column:market_price"`
		ShopPrice       float32  `gorm:"column:shop_price"`
		GoodsBrief      string   `gorm:"column:goods_brief"`
		ShipFree        bool     `gorm:"column:ship_free"`
		Images          GormList `gorm:"column:images"`
		DescImages      GormList `gorm:"column:desc_images"`
		GoodsFrontImage string   `gorm:"column:goods_front_image"`
		IsNew           bool     `gorm:"column:is_new"`
		IsHot           bool     `gorm:"column:is_hot"`
	}
)

//
//func (g *Goods) AfterCreate(tx *gorm.DB) (err error) {
//	Esmodel := EsGoods{
//		ID:          g.ID,
//		CategoryID:  g.CategoryID,
//		BrandsID:    g.BrandsID,
//		OnSale:      g.OnSale,
//		ShipFree:    g.ShipFree,
//		IsNew:       g.IsNew,
//		IsHot:       g.IsHot,
//		Name:        g.Name,
//		ClickNum:    g.ClickNum,
//		SoldNum:     g.SoldNum,
//		FavNum:      g.FavNum,
//		MarketPrice: g.MarketPrice,
//		GoodsBrief:  g.GoodsBrief,
//		ShopPrice:   g.ShopPrice,
//	}
//
//	_, err = global.Esclient.Index().Index(EsGoods{}.GetindexName()).BodyJson(Esmodel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (g *Goods) AfterUpdate(tx *gorm.DB) (err error) {
//	Esmodel := EsGoods{
//		ID:          g.ID,
//		CategoryID:  g.CategoryID,
//		BrandsID:    g.BrandsID,
//		OnSale:      g.OnSale,
//		ShipFree:    g.ShipFree,
//		IsNew:       g.IsNew,
//		IsHot:       g.IsHot,
//		Name:        g.Name,
//		ClickNum:    g.ClickNum,
//		SoldNum:     g.SoldNum,
//		FavNum:      g.FavNum,
//		MarketPrice: g.MarketPrice,
//		GoodsBrief:  g.GoodsBrief,
//		ShopPrice:   g.ShopPrice,
//	}
//
//	_, err = global.Esclient.Update().Index(EsGoods{}.GetindexName()).Doc(Esmodel).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (g *Goods) AfterDelete(tx *gorm.DB) (err error) {
//
//	_, err = global.Esclient.Delete().Index(EsGoods{}.GetindexName()).Id(strconv.Itoa(int(g.ID))).Do(context.Background())
//	if err != nil {
//		return err
//	}
//	return nil
//}

func newGoodsModel(conn *gorm.DB, c cache.CacheConf) *defaultGoodsModel {
	return &defaultGoodsModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`goods`",
	}
}

func (m *defaultGoodsModel) Insert(ctx context.Context, data *Goods) error {
	goZeroMallGoodsIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallGoodsIdPrefix, data.Id)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Save(&data).Error
	}, goZeroMallGoodsIdKey)
	return err
}

func (m *defaultGoodsModel) FindOne(ctx context.Context, id int64) (*Goods, error) {
	goZeroMallGoodsIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallGoodsIdPrefix, id)
	var resp Goods
	err := m.QueryCtx(ctx, &resp, goZeroMallGoodsIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Goods{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultGoodsModel) Update(ctx context.Context, data *Goods) error {
	goZeroMallGoodsIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallGoodsIdPrefix, data.Id)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Save(data).Error
	}, goZeroMallGoodsIdKey)
	return err
}

func (m *defaultGoodsModel) Delete(ctx context.Context, id int64) error {
	goZeroMallGoodsIdKey := fmt.Sprintf("%s%v", cacheGoZeroMallGoodsIdPrefix, id)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Delete(&Goods{}, id).Error
	}, goZeroMallGoodsIdKey)
	return err
}

func (m *defaultGoodsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGoZeroMallGoodsIdPrefix, primary)
}

func (m *defaultGoodsModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&Goods{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultGoodsModel) tableName() string {
	return m.table
}