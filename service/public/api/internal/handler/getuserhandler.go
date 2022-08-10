package handler

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	errorx "zero-mal/common/error"
	"zero-mal/common/initialize"
	"zero-mal/global"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-mal/service/public/api/internal/logic"
	"zero-mal/service/public/api/internal/svc"
	"zero-mal/service/public/api/internal/types"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		//增加个翻译器.没有生效，需要去检查下
		if err := global.Validate.Struct(&req); err != nil {
			errs, _ := err.(validator.ValidationErrors)
			m := errs.Translate(global.Trans)
			errNew := initialize.ReturnString(initialize.RemoveTopStruct(m))
			httpx.Error(w, errorx.NewDefaultError(errNew))
			return
		}

		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.Error(w, errorx.NewDefaultError(err.Error()))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
