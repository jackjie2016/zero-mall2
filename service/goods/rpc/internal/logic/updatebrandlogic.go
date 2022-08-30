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

type UpdateBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandLogic {
	return &UpdateBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateBrandLogic) UpdateBrand(in *goods_pb.BrandRequest) (*goods_pb.Empty, error) {
	// todo: add your logic here and delete this line
	var brand *model.Brands
	var err error

	if brand, err = l.svcCtx.BrandsModel.FindOne(l.ctx, int64(in.Id)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	if len(in.Name) > 0 {
		brand.Name = in.Name
	}

	if len(in.Logo) > 0 {
		brand.Logo = in.Logo
	}

	if err := l.svcCtx.BrandsModel.Update(l.ctx, brand); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "品牌更新失败")
	}
	return &goods_pb.Empty{}, nil
}
