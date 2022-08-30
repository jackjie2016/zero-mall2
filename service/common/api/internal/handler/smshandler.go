package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	errorx "zero-mal/common/error"
	"zero-mal/common/initialize"
	"zero-mal/global"
	"zero-mal/service/common/api/internal/logic"
	"zero-mal/service/common/api/internal/svc"
	"zero-mal/service/common/api/internal/types"
)

func SmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SmsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)

			return
		}
		////增加个翻译器.没有生效，需要去检查下
		if err := global.Validate.Struct(&req); err != nil {
			errs, _ := err.(validator.ValidationErrors)
			m := errs.Translate(global.Trans)
			errNew := initialize.ReturnString(initialize.RemoveTopStruct(m))
			httpx.Error(w, errorx.NewDefaultError(errNew))
			return
		}
		//增加个翻译器
		//if err := validator.New().StructCtx(r.Context(), &req); err != nil {
		//	httpx.Error(w, errorx.NewDefaultError(err.Error()))
		//	return
		//}
		l := logic.NewSmsLogic(r.Context(), svcCtx)
		resp, err := l.Sms(&req)
		if err != nil {
			httpx.Error(w, errorx.NewDefaultError(err.Error()))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
