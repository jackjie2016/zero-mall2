package logic

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"strings"
	"time"
	errorx "zero-mal/common/error"

	"zero-mal/service/common/api/internal/svc"
	"zero-mal/service/common/api/internal/types"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/zeromicro/go-zero/core/logx"
)

type SmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmsLogic {
	return &SmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func GenerateSmsCode(witdh int) string {
	//生成width长度的短信验证码

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < witdh; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
func (l *SmsLogic) Sms(req *types.SmsRequest) (resp *types.SmsResponse, err error) {
	// todo: add your logic here and delete this line

	mobile := req.Mobile
	var code string

	l.svcCtx.Cache.Get(fmt.Sprintf("sms:%s", mobile), &code)
	if len(code) > 0 {
		return nil, errorx.NewDefaultError("已发送，重复点击")
	}
	return nil, errorx.NewDefaultError("已发送，重复点击")
	code = GenerateSmsCode(6)
	AliSmsInfo := l.svcCtx.Config.AliSmsInfo
	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", AliSmsInfo.ApiKey, AliSmsInfo.ApiSecret)
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = mobile                     //手机号
	request.QueryParams["SignName"] = AliSmsInfo.SignName            //阿里云验证过的项目名 自己设置
	request.QueryParams["TemplateCode"] = AliSmsInfo.TemplateCode    //阿里云的短信模板号 自己设置
	request.QueryParams["TemplateParam"] = "{\"code\":" + code + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	response, err := client.ProcessCommonRequest(request)

	zap.S().Infof("%v", client.DoAction(request, response))

	if err != nil {
		return nil, errorx.NewDefaultError("发送失败")
	}
	l.svcCtx.Cache.SetWithExpire(fmt.Sprintf("sms:%s", mobile), code, time.Duration(time.Second*60))
	//mobile, code
	return &types.SmsResponse{}, nil
}
