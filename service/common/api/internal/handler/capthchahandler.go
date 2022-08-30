package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	errorx "zero-mal/common/error"
	"zero-mal/service/common/api/internal/logic"
	"zero-mal/service/common/api/internal/svc"
)

func CapthchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCapthchaLogic(r.Context(), svcCtx)
		resp, err := l.Capthcha()
		if err != nil {
			httpx.Error(w, errorx.NewDefaultError(err.Error()))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
