package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/service/goods/rpc/goods_pb"
	"zero-mal/service/goods/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(in *goods_pb.DeleteCategoryRequest) (*goods_pb.Empty, error) {
	// todo: add your logic here and delete this line

	if err := l.svcCtx.CategoryModel.Delete(l.ctx, int64(in.Id)); err != nil {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	return &goods_pb.Empty{}, nil
}
