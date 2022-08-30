package user

import (
	"context"
	"zero-mal/service/user/api/internal/svc"
	"zero-mal/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.Createrequest) error {
	// todo: add your logic here and delete this line
	////InitGrpc()
	//RegisterForm := forms.RegisterForm{}
	//if err := c.ShouldBind(&RegisterForm); err != nil {
	//	HandleValitor(c, err)
	//	return
	//}
	//
	////验证码
	////rdb := redis.NewClient(&redis.Options{
	////	Addr:     fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	////	Password: "", // no password set
	////	DB:       0,  // use default DB
	////})
	//rdb := redis.NewClient(&redis.Options{
	//	Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	//})
	//val2, err := rdb.Get(context.Background(), RegisterForm.Mobile).Result()
	//if err == redis.Nil {
	//	zap.S().Errorf("Redis中的验证码错误: %s", val2)
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": "验证码错误",
	//	})
	//	return
	//} else {
	//	zap.S().Infof("Redis中的验证码: %s", val2)
	//	if val2 != RegisterForm.Code {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"code": "验证码错误",
	//		})
	//		return
	//	}
	//}
	//
	////注册的逻辑 查询手机是否存在
	//if _, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
	//	Mobile: RegisterForm.Mobile,
	//}); err == nil {
	//	c.JSON(http.StatusConflict, map[string]string{
	//		"msg": "手机号已经被注册",
	//	})
	//	return
	//} else {
	//	if e, ok := status.FromError(err); ok {
	//		fmt.Println(e.Code())
	//	}
	//}
	//
	////只是查询到用户了而已，并没有检查密码
	//if regRsp, pasErr := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
	//	Mobile:   RegisterForm.Mobile,
	//	NickName: RegisterForm.Mobile,
	//	Password: RegisterForm.PassWord,
	//}); pasErr != nil {
	//	zap.S().Errorf("[Register] 查询 【新建用户失败】失败: %s", err.Error())
	//	HandleGrpcErrorToHttp(err, c)
	//	return
	//} else {
	//	//生成token
	//	j := middlewares.NewJWT()
	//	claim := models.CustomClaims{
	//		ID:          uint(regRsp.Id),
	//		NickName:    regRsp.NickName,
	//		AuthorityID: uint(regRsp.Role),
	//		StandardClaims: jwt.StandardClaims{
	//			Audience:  "",                              //观众
	//			ExpiresAt: time.Now().Unix() + 30*60*60*24, //一个月
	//			Issuer:    "zifeng6257",                    //观众
	//			NotBefore: time.Now().Unix(),
	//			Subject:   "hhaha",
	//		},
	//	}
	//	var token string
	//	if token, err = j.CreateToken(claim); err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"msg": "生成token失败",
	//		})
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"id":         regRsp.Id,
	//		"nick_name":  regRsp.NickName,
	//		"token":      token,
	//		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	//	})
	//
	//}

	return nil
}
