// Code generated by goctl. DO NOT EDIT.
package types

type SmsRequest struct {
	Mobile string `form:"mobile"  validate:"required,min=3,max=11,mobile"`
}

type SmsResponse struct {
}

type CapthchaRequest struct {
}

type CapthchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}
