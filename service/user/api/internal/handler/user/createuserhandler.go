package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-mal/service/user/api/internal/logic/user"
	"zero-mal/service/user/api/internal/svc"
	"zero-mal/service/user/api/internal/types"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Createrequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewCreateUserLogic(r.Context(), svcCtx)
		err := l.CreateUser(&req)
		if err != nil {

			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
