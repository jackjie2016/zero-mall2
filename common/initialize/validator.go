package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	myvalidator "zero-mal/common/validator"
	"zero-mal/global"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func InitTrans(locale string) (err error) {
	global.Validate = validator.New()
	//修改gin框架中的validator 引擎属性，实现定制
	global.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	zhT := zh.New() //中文翻译器
	enT := en.New() //英文翻译器
	//第一个参数是备用的语言环境，后面的应该支持的语言环境
	uni := ut.New(enT, zhT, zhT)
	var ok bool
	global.Trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s)", locale)
	}
	switch locale {
	case "en":
		en_translations.RegisterDefaultTranslations(global.Validate, global.Trans)
	case "zh":
		zh_translations.RegisterDefaultTranslations(global.Validate, global.Trans)
	default:
		en_translations.RegisterDefaultTranslations(global.Validate, global.Trans)
	}

	_ = global.Validate.RegisterValidation("mobile", myvalidator.ValidateMobile)
	_ = global.Validate.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})

	return
}

func RemoveTopStruct(fields map[string]string) map[string]string {

	rsp := map[string]string{}
	for field, err := range fields {
		//fmt.Printf("field:[%s],err:[%s]\n", field, err)
		//
		//fmt.Printf("位置:[%s]\n", field[strings.Index(field, ".")+1:])
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func ReturnString(fields map[string]string) string {
	msgs := RemoveTopStruct(fields)
	return MapToJson(msgs)
}

func MapToJson(m map[string]string) string {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(jsonByte)
}
