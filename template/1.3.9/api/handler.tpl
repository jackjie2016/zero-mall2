package {{.PkgName}}

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
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

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			httpx.Error(w, errorx.NewDefaultError(err.Error()))
		} else {
			{{if .HasResp}}httpx.OkJson(w, resp){{else}}httpx.Ok(w){{end}}
		}
	}
}
