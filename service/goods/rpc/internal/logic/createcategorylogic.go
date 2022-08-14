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

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryLogic) CreateCategory(in *pb.CategoryInfoRequest) (*pb.CategoryInfoResponse, error) {
	// todo: add your logic here and delete this line
	category := model.Category{}

	category.Name = in.Name
	category.Level = in.Level
	if in.Level != 1 {
		//去查询父类目是否存在
		category.ParentCategoryID = in.ParentCategory
	}
	category.IsTab = in.IsTab

	//global.DB.Save(&category)

	if err := l.svcCtx.CategoryModel.Insert(l.ctx, &category); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "分类创建失败")
	}

	return &pb.CategoryInfoResponse{
		Id:             category.Id,
		Name:           category.Name,
		ParentCategory: category.ParentCategoryID,
		Level:          category.Level,
		IsTab:          category.IsTab,
	}, nil

}
