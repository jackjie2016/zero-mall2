package user

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	errorx "zero-mal/common/error"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-mal/service/user/api/internal/logic/user"
	"zero-mal/service/user/api/internal/svc"
	"zero-mal/service/user/api/internal/types"
)

func GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Inforequest
		//logx.Error("测试Error的日志")
		//logx.Info("测试Info的日志")
		//logx.Slow("测试Slow的日志")
		//logx.Stat("测试Stat的日志")
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(err.Error()))
			return
		}

		//增加个翻译器
		if err := validator.New().StructCtx(r.Context(), &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError(err.Error()))
			return
		}

		l := user.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
