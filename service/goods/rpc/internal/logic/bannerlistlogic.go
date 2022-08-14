package logic

import (
	"context"
	"zero-mal/global"
	model "zero-mal/service/goods/model/gorm"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

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
func (l *BannerListLogic) BannerList(in *pb.Empty) (*pb.BannerListResponse, error) {
	// todo: add your logic here and delete this line
	bannerListResponse := pb.BannerListResponse{}

	var banners []model.Banner
	result := global.DB.Find(&banners)
	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerReponses []*pb.BannerResponse
	for _, banner := range banners {
		bannerReponses = append(bannerReponses, &pb.BannerResponse{
			Id:    banner.Id,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		})
	}

	bannerListResponse.Data = bannerReponses

	return &bannerListResponse, nil

}
