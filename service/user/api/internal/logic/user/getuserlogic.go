package user

import (
	"context"
	"fmt"
	errorx "zero-mal/common/error"
	"zero-mal/service/user/api/internal/svc"
	"zero-mal/service/user/api/internal/types"
	pb "zero-mal/service/user/rpc/user_pb"

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

	response, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &pb.GetUserByIdReq{Id: req.Id})
	fmt.Println(response)

	if err != nil {
		return nil, errorx.NewDefaultError("用户不存在")
	}

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
		Desc:     response.User.Desc,
		HeadUrl:  response.User.HeadUrl,
	}, nil
}
