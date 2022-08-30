package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/common/tool"
	"zero-mal/global"
	esmodel "zero-mal/service/goods/model/es"
	model "zero-mal/service/goods/model/gorm"
	"zero-mal/service/goods/rpc/goods_pb"
	"zero-mal/service/goods/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品接口
func (l *GoodsListLogic) GoodsList(in *goods_pb.GoodsFilterRequest) (*goods_pb.GoodsListResponse, error) {
	// todo: add your logic here and delete this line

	//l.svcCtx.GoodsModel.FindOne()

	//关键词搜索、查询新品、查询热门商品、通过价格区间筛选， 通过商品分类筛选
	goodsListResponse := &goods_pb.GoodsListResponse{}
	localDB := global.DB.Model(&model.Goods{})

	q := elastic.NewBoolQuery()

	var goods []model.Goods

	if in.KeyWords != "" {
		//localDB = localDB.Where("name Like ?", "%"+in.KeyWords+"%")
		//多字段查询
		q = q.Must(elastic.NewMultiMatchQuery(in.KeyWords, "name", "goods_brief"))
	}
	if in.IsHot {
		//localDB = localDB.Where(model.Goods{IsHot: true})
		q = q.Filter(elastic.NewTermQuery("is_hot", in.IsHot))
	}
	if in.IsNew {
		//localDB = localDB.Where(model.Goods{IsNew: true})
		q = q.Filter(elastic.NewTermQuery("is_new", in.IsNew))
	}

	if in.PriceMin > 0 {
		//localDB = localDB.Where("shop_price>=?", in.PriceMin)
		q = q.Filter(elastic.NewRangeQuery("shop_price").Gte(in.PriceMin))
	}

	if in.PriceMax > 0 {
		q = q.Filter(elastic.NewRangeQuery("shop_price").Lte(in.PriceMax))
	}

	if in.TopCategory > 0 {
		//通过category去查询商品
		var subQuery string

		var category model.Category
		if result := global.DB.First(&category, in.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}

		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category WHERE parent_category_id=%d)", in.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category WHERE parent_category_id=%d", in.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from category WHERE id=%d", in.TopCategory)
		}
		categoryIDs := make([]interface{}, 0)

		type CateResult struct {
			ID int32
		}
		var CateResults []CateResult
		global.DB.Model(model.Category{}).Raw(subQuery).Scan(&CateResults)
		for _, c := range CateResults {
			categoryIDs = append(categoryIDs, c.ID)
		}

		//localDB = localDB.Where("category_id in ?", categoryIDs)
		//es 中根据分类id过滤
		q = q.Filter(elastic.NewTermsQuery("category_id", categoryIDs...))
	}
	if in.Pages == 0 {
		in.Pages = 1
	}

	switch {
	case in.PagePerNums > 100:
		in.PagePerNums = 100
	case in.PagePerNums <= 0:
		in.PagePerNums = 10
	}

	goodsIds := make([]int32, 0)
	esRes, err := global.Esclient.Search().Index(esmodel.EsGoods{}.GetindexName()).Query(q).From(int(in.Pages)).Size(int(in.PagePerNums)).Do(context.Background())
	if err != nil {
		return nil, err
	}

	goodsListResponse.Total = int32(esRes.Hits.TotalHits.Value)

	if goodsListResponse.Total == 0 {
		return nil, status.Errorf(codes.ResourceExhausted, "没有查询到商品")
	}

	for _, value := range esRes.Hits.Hits {
		goods := esmodel.EsGoods{}
		_ = json.Unmarshal(value.Source, &goods)
		goodsIds = append(goodsIds, goods.ID)
	}

	result := localDB.Preload("Category").Preload("Brands").Find(&goods, goodsIds)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	return goodsListResponse, nil
}

// 商品接口
func (l *GoodsListLogic) GoodsList2(in *goods_pb.GoodsFilterRequest) (*goods_pb.GoodsListResponse, error) {
	// todo: add your logic here and delete this line

	//l.svcCtx.GoodsModel.FindOne()

	//关键词搜索、查询新品、查询热门商品、通过价格区间筛选， 通过商品分类筛选
	goodsListResponse := &goods_pb.GoodsListResponse{}
	localDB := global.DB.Model(&model.Goods{})

	//q := elastic.NewBoolQuery()
	//q = q.Must(NewTermQuery("tag", "wow"))
	//q = q.Filter(NewTermQuery("account", "1"))

	var goods []model.Goods

	if in.KeyWords != "" {
		localDB = localDB.Where("name Like ?", "%"+in.KeyWords+"%")
		//多字段查询
		//q = q.Must(elastic.NewMultiMatchQuery(req.KeyWords, "name", "goods_brief"))
	}
	if in.IsHot {
		localDB = localDB.Where(model.Goods{IsHot: true})
		//q = q.Filter(elastic.NewTermQuery("is_hot", req.IsHot))
	}
	if in.IsNew {
		localDB = localDB.Where(model.Goods{IsNew: true})
		//q = q.Filter(elastic.NewTermQuery("is_new", req.IsNew))
	}

	if in.PriceMin > 0 {
		localDB = localDB.Where("shop_price>=?", in.PriceMin)
		//q = q.Filter(elastic.NewRangeQuery("shop_price").Gte(req.PriceMin))
	}

	if in.PriceMax > 0 {
		localDB = localDB.Where("shop_price <=?", in.PriceMax)
		//q = q.Filter(elastic.NewRangeQuery("shop_price").Lte(req.PriceMax))
	}

	if in.TopCategory > 0 {
		//通过category去查询商品
		var subQuery string

		var category model.Category
		if result := global.DB.First(&category, in.TopCategory); result.RowsAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}

		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category WHERE parent_category_id=%d)", in.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category WHERE parent_category_id=%d", in.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from category WHERE id=%d", in.TopCategory)
		}
		categoryIDs := make([]interface{}, 0)

		type CateResult struct {
			ID int32
		}
		var CateResults []CateResult
		//global.DB.Where(fmt.Sprintf("category_id in (%s)", subQuery)).Find(CateResults)
		global.DB.Model(model.Category{}).Raw(subQuery).Scan(&CateResults)
		for _, c := range CateResults {
			categoryIDs = append(categoryIDs, c.ID)
		}

		localDB = localDB.Where("category_id in ?", categoryIDs)
		//es 中根据分类id过滤
		//q = q.Filter(elastic.NewTermsQuery("category_id", categoryIDs...))
	}
	if in.Pages == 0 {
		in.Pages = 1
	}

	switch {
	case in.PagePerNums > 100:
		in.PagePerNums = 100
	case in.PagePerNums <= 0:
		in.PagePerNums = 10
	}

	resultp := localDB.Find(&goods)
	if resultp.Error != nil {
		return nil, resultp.Error
	}
	//goodsListResponse.Total = int32(esRes.Hits.TotalHits.Value)

	//if goodsListResponse.Total == 0 {
	//	return nil, status.Errorf(codes.ResourceExhausted, "没有查询到商品")
	//}
	goodsListResponse.Total = int32(resultp.RowsAffected)
	//for _, value := range esRes.Hits.Hits {
	//	goods := model.EsGoods{}
	//	_ = json.Unmarshal(value.Source, &goods)
	//	goodsIds = append(goodsIds, goods.ID)
	//}

	//var count int64
	//localDB.Count(&count)
	//goodsListResponse.Total = int32(count)

	result := localDB.Preload("Category").Preload("Brands").Scopes(tool.Paginate(int(in.Pages), int(in.PagePerNums))).Find(&goods)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, good := range goods {
		//var goodsInfoResponse goods_pb.GoodsInfoResponse
		//_ = copier.Copy(&goodsInfoResponse, good)// 出现categoryId值不能都是0的情况

		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	return goodsListResponse, nil
}
func ModelToResponse(goods model.Goods) goods_pb.GoodsInfoResponse {
	return goods_pb.GoodsInfoResponse{
		Id:              goods.Id,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsBrief:      goods.GoodsBrief,
		ShipFree:        goods.ShipFree,
		GoodsFrontImage: goods.GoodsFrontImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		DescImages:      goods.DescImages,
		Images:          goods.Images,
		Category: &goods_pb.CategoryBriefInfoResponse{
			Id:   goods.Category.Id,
			Name: goods.Category.Name,
		},
		Brand: &goods_pb.BrandInfoResponse{
			Id:   goods.Brands.Id,
			Name: goods.Brands.Name,
			Logo: goods.Brands.Logo,
		},
	}
}
