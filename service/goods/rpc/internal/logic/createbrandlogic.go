package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/global"
	model "zero-mal/service/goods/model/gorm"

	"zero-mal/service/goods/rpc/goods_pb"
	"zero-mal/service/goods/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBrandLogic {
	return &CreateBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateBrandLogic) CreateBrand(in *goods_pb.BrandRequest) (*goods_pb.BrandInfoResponse, error) {
	// todo: add your logic here and delete this line

	//新建品牌
	if result := global.DB.Where("name=?", in.Name).First(&model.Brands{}); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}
	brands := &model.Brands{
		Name: in.Name,
		Logo: in.Logo,
	}

	global.DB.Save(brands)
	return &goods_pb.BrandInfoResponse{Id: brands.Id, Name: brands.Name, Logo: brands.Logo}, nil
}
