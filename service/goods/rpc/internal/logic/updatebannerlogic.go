package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	model "zero-mal/service/goods/model/gorm"
	"zero-mal/service/goods/rpc/goods_pb"
	"zero-mal/service/goods/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBannerLogic {
	return &UpdateBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBannerLogic) UpdateBanner(in *goods_pb.BannerRequest) (*goods_pb.Empty, error) {
	// todo: add your logic here and delete this line
	//banner := model.Banner{
	//	BaseModel: model.BaseModel{Id: in.Id},
	//	Image:     in.Image,
	//	Url:       in.Url,
	//	Index:     in.Index,
	//}
	var banner *model.Banner
	var err error
	if banner, err = l.svcCtx.BannerModel.FindOne(l.ctx, int64(in.Id)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "广告不存在")
	}

	if err := l.svcCtx.BannerModel.Update(l.ctx, banner); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "广告更新失败")
	}
	return &goods_pb.Empty{}, nil

}
