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

type UpdateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryLogic {
	return &UpdateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCategoryLogic) UpdateCategory(in *pb.CategoryInfoRequest) (*pb.Empty, error) {
	// todo: add your logic here and delete this line
	var category *model.Category

	var err error
	if category, err = l.svcCtx.CategoryModel.FindOne(l.ctx, int64(in.Id)); err != nil {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	if in.Name != "" {
		category.Name = in.Name
	}
	if in.ParentCategory != 0 {
		category.ParentCategoryID = in.ParentCategory
	}
	if in.Level != 0 {
		category.Level = in.Level
	}
	if in.IsTab {
		category.IsTab = in.IsTab
	}
	if err = l.svcCtx.CategoryModel.Update(l.ctx, category); err != nil {
		return nil, status.Errorf(codes.Internal, "更新失败")
	}

	return &pb.Empty{}, nil
}
