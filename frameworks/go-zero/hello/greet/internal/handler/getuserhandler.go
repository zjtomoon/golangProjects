package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"golangProjects/frameworks/go-zero/hello/greet/internal/logic"
	"golangProjects/frameworks/go-zero/hello/greet/internal/svc"
	"golangProjects/frameworks/go-zero/hello/greet/internal/types"
)

func GetUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), ctx)
		resp, err := l.GetUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
