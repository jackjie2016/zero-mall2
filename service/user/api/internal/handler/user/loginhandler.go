package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	errorx "zero-mal/common/error"
	"zero-mal/common/initialize"
	"zero-mal/global"
	"zero-mal/service/user/api/internal/logic/user"
	"zero-mal/service/user/api/internal/svc"
	"zero-mal/service/user/api/internal/types"
)

var validate *validator.Validate

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Loginrequest

		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(err.Error()))
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

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
