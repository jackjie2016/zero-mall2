package logic

import (
	"context"
	"fmt"
	errorx "zero-mal/common/error"
	"zero-mal/common/tool"

	"github.com/mojocn/base64Captcha"

	"zero-mal/service/common/api/internal/svc"
	"zero-mal/service/common/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CapthchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCapthchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CapthchaLogic {
	return &CapthchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CapthchaLogic) Capthcha() (resp *types.CapthchaResponse, err error) {
	// todo: add your logic here and delete this line
	var store = tool.RedisStore{}
	fmt.Println(222222)
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)

	cp := base64Captcha.NewCaptcha(driver, store)
	fmt.Println(cp)
	id, b64s, err := cp.Generate()
	if err != nil {
		logx.Errorf("生成验证码错误,: ", err.Error())
		return nil, errorx.NewDefaultError("生成验证码错误")
	}

	return &types.CapthchaResponse{CaptchaId: id, PicPath: b64s}, nil
}
