package logic

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"zero-mal/global"
	model "zero-mal/service/user/model/genModel"
	"zero-mal/service/user/rpc/internal/svc"
	"zero-mal/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	// todo: add your logic here and delete this line
	var user model.User
	result := global.DB.First(&user, in.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.AlreadyExists, "用户不存在")
	}
	var Birthday time.Time
	if in.Birthday > 0 {
		Birthday = time.Unix(int64(in.Birthday), 0)
		user.Birthday = &Birthday
	}
	if in.NickName != "" {
		user.NickName = in.NickName
	}
	if in.Gender != "" {
		user.Gender = in.Gender
	}
	if in.HeadUrl != "" {
		user.HeadUrl = in.HeadUrl
	}
	if in.Address != "" {
		user.Address = in.Address
	}
	if in.Mobile != "" {
		user.Mobile = in.Mobile
	}
	if in.Desc != "" {
		user.Desc = in.Desc
	}
	if in.Role > -1 {
		user.Role = in.Role
	}
	//1 、修改密码
	if in.Password != "" {
		options := &password.Options{6, 100, 30, sha512.New}
		salt, encodedPwd := password.Encode(in.Password, options)
		user.Password = fmt.Sprintf("$zifeng-sha512$%s$%s", salt, encodedPwd)
	}

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	var respUser pb.UpdateUserResp
	_ = copier.Copy(&respUser, user)
	return &respUser, nil
}
