package logic

import (
	"context"
	"zero-mal/common/tool"
	"zero-mal/global"
	model "zero-mal/service/goods/model/gorm"

	"zero-mal/service/goods/rpc/goods_pb"
	"zero-mal/service/goods/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandListLogic {
	return &BrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 品牌和轮播图
func (l *BrandListLogic) BrandList(in *goods_pb.BrandFilterRequest) (*goods_pb.BrandListResponse, error) {
	// todo: add your logic here and delete this line

	var brands []model.Brands
	result := global.DB.Find(&brands)
	//查询没有错误
	if result.Error != nil {
		return nil, result.Error
	}
	brandListResponse := &goods_pb.BrandListResponse{}
	brandListResponse.Total = int32(result.RowsAffected)
	global.DB.Scopes(tool.Paginate(int(in.Pages), int(in.PagePerNums))).Find(&brands)

	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	brandListResponse.Total = int32(total)

	for _, brand := range brands {
		brandListResponse.Data = append(brandListResponse.Data, &goods_pb.BrandInfoResponse{
			Id:   brand.Id,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	return brandListResponse, nil
}
