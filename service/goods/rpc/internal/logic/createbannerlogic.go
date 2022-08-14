package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	model "zero-mal/service/goods/model/gorm"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBannerLogic {
	return &CreateBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBannerLogic) CreateBanner(in *pb.BannerRequest) (*pb.BannerResponse, error) {
	// todo: add your logic here and delete this line

	banner := model.Banner{
		Image: in.Image,
		Url:   in.Url,
		Index: in.Index,
	}

	//新建品牌
	if err := l.svcCtx.BannerModel.Insert(l.ctx, &banner); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "广告创建失败")
	}
	var respose pb.BannerResponse
	respose.Id = banner.Id
	respose.Image = banner.Image
	respose.Index = banner.Index
	respose.Url = banner.Url
	return &respose, nil
}
