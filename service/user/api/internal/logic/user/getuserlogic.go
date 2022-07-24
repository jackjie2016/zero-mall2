package user

import (
	"context"
	"fmt"
	"zero-mal/service/user/api/internal/svc"
	"zero-mal/service/user/api/internal/types"
	"zero-mal/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.Inforequest) (resp *types.Inforesponse, err error) {
	// todo: add your logic here and delete this line

	response, _ := l.svcCtx.UserRpc.GetUserById(l.ctx, &pb.GetUserByIdReq{Id: req.Id})
	fmt.Println(response)

	//msg, _ := json.Marshal(struct {
	//	Code int64
	//	Msg  string
	//}{Code: 200, Msg: "验证成功"})
	//w.Write(msg)
	//
	//return nil, errorx.NewDefaultError("用户不存在")

	return &types.Inforesponse{
		ID:       req.Id,
		Username: response.User.Mobile,
		NickName: response.User.NickName,
		Gender:   response.User.Gender,
	}, nil
}
