package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"golangProjects/frameworks/web/mvc/go-zero/hello/greet/internal/logic"
	"golangProjects/frameworks/web/mvc/go-zero/hello/greet/internal/svc"
	"golangProjects/frameworks/web/mvc/go-zero/hello/greet/internal/types"
)

func GreetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), svcCtx)
		resp, err := l.Greet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
