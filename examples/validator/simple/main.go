package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"zero-mal/common/initialize"
	"zero-mal/global"
)

func main() {

	//4、初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err.Error())
	}

	validateStruct()

}

func validateStruct() {
	type Loginrequest struct {
		Mobile    string `form:"mobile" json:"mobile" validate:"required,min=3,max=11,mobile"`
		Password  string `form:"password" json:"password" validate:"required,min=5,max=10"`
		Captcha   string `form:"captcha" json:"captcha" validate:"required,min=5,max=5"`
		CaptchaId string `form:"captcha_id" json:"captchaId" validate:"required"`
	}
	Login := &Loginrequest{
		Mobile:    "15958615799",
		Password:  "1555",
		Captcha:   "1555",
		CaptchaId: "1555",
	}
	err := global.Validate.Struct(Login)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		errs, _ := err.(validator.ValidationErrors)
		m := errs.Translate(global.Trans)
		errNew := initialize.RemoveTopStruct(m)

		fmt.Println(errNew)

		// from here you can create your own error messages in whatever language you wish
		return
	}
}
