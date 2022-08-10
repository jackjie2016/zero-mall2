package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero-mal/common/xerr"
	Grommodel "zero-mal/service/user/model/gorm"

	"zero-mal/service/user/rpc/internal/svc"
	"zero-mal/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserLogic {
	return &DelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserLogic) DelUser(in *pb.DelUserReq) (*pb.DelUserResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserGormModel.FindOne(l.ctx, in.Id)
	if err != nil && err != Grommodel.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user  fail"), "err : %v , in : %+v", err, in)
	}
	err = l.svcCtx.UserGormModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除失败")
	}

	return &pb.DelUserResp{}, nil
}
