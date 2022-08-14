package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBrandLogic {
	return &DeleteBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBrandLogic) DeleteBrand(in *pb.BrandRequest) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	if err := l.svcCtx.BrandsModel.Delete(l.ctx, int64(in.Id)); err != nil {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &pb.Empty{}, nil
}
