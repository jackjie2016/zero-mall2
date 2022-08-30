package logic

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"zero-mal/service/user/rpc/internal/svc"
	pb "zero-mal/service/user/rpc/user_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckPassWordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckPassWordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPassWordLogic {
	return &CheckPassWordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckPassWordLogic) CheckPassWord(in *pb.CheckInfo) (*pb.CheckResponse, error) {
	// todo: add your logic here and delete this line
	options := &password.Options{6, 100, 30, sha512.New}
	passwordinfo := strings.Split(in.EncryptedPassword, "$")
	fmt.Println(len(passwordinfo))
	if len(passwordinfo) != 4 {
		return nil, status.Errorf(codes.InvalidArgument, "加密密码不合法")
	}
	check := password.Verify(in.Password, passwordinfo[2], passwordinfo[3], options)
	//TransCtx(ctx context.Context, session sqlx.Session) error
	//err := l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
	//	//3、构建用户模型
	//
	//	toBeCharge := "1988-01-01 00:00:00" //待转化为时间戳的字符串 注意 这里的小时
	//	//转化所需模板
	//	loc, _ := time.LoadLocation("Local")                                       //获取时区
	//	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	//
	//	userData := &genModel.User{
	//		//Id:       0,
	//		Mobile:   "111111",
	//		Password: "2222",
	//		NickName: "",
	//		HeadUrl:  "",
	//		Birthday: &theTime,
	//		Address:  "",
	//		Desc:     "",
	//		Gender:   0,
	//		Role:     0,
	//	}
	//	response, err := l.svcCtx.UserModel.Insert(l.ctx, userData)
	//
	//	fmt.Println(response, err)
	//	return nil
	//})
	//fmt.Println(err)
	return &pb.CheckResponse{Success: check}, nil
}
