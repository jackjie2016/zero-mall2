package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/service/goods/rpc/internal/svc"
	"zero-mal/service/goods/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBannerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBannerLogic {
	return &DeleteBannerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteBannerLogic) DeleteBanner(in *pb.BannerRequest) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	fmt.Println("ok")
	if err := l.svcCtx.BannerModel.Delete(l.ctx, int64(in.Id)); err != nil {
		return nil, status.Errorf(codes.NotFound, "广告不存在")
	}
	return &pb.Empty{}, nil
}
