package logic

import (
	"context"
	"zero-mal/global"
	model "zero-mal/service/goods/model/gorm"

	"zero-mal/service/goods/rpc/goods_pb"
	"zero-mal/service/goods/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BannerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BannerListLogic {
	return &BannerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 轮播图
func (l *BannerListLogic) BannerList(in *goods_pb.Empty) (*goods_pb.BannerListResponse, error) {
	// todo: add your logic here and delete this line
	bannerListResponse := goods_pb.BannerListResponse{}

	var banners []model.Banner
	result := global.DB.Find(&banners)
	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerReponses []*goods_pb.BannerResponse
	for _, banner := range banners {
		bannerReponses = append(bannerReponses, &goods_pb.BannerResponse{
			Id:    banner.Id,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}

	bannerListResponse.Data = bannerReponses

	return &bannerListResponse, nil

}
