# 数据库生成

gentool -dsn "ai-chat:E6zhRdc3J8L6c4hn@tcp(47.113.177.135:3306)/ai-chat?charset=utf8mb4&parseTime=True&loc=Local" -db "mysql" -outPath "../internal/infra/mysql/query" -outFile "gen.go" -fieldWithIndexTag true -fieldWithTypeTag true -modelPkgName "mysql"

# api生成模板

goctl api go -api api.api -dir ../  --style=goZero

goctl template init

vim ~/.goctl/${goctl版本号}/api/handler.tpl
handler.tpl

package {{.PkgName}}

import (
xhttp "github.com/zeromicro/x/http"
"net/http"

	{{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
	{{.ImportPackages}}

)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
{{if .HasRequest}}var req types.{{.RequestType}}
if err := httpx.Parse(r, &req); err != nil {
httpx.ErrorCtx(r.Context(), w, err)
return
}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		{{if .HasResp}}
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
		{{else}}xhttp.JsonBaseResponseCtx(r.Context(), w, err){{end}}
	}

}

goctl api go -api *.api -dir ../  --style=goZero    
docker-compose up -d --build order-center