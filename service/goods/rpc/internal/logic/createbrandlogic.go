package logic

import (
	"context"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

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

func (l *CreateBrandLogic) CreateBrand(in *pb.BrandRequest) (*pb.BrandInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.BrandInfoResponse{}, nil
}
