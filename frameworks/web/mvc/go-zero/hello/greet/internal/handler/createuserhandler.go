package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"golangProjects/frameworks/go-zero/hello/greet/internal/logic"
	"golangProjects/frameworks/go-zero/hello/greet/internal/svc"
	"golangProjects/frameworks/go-zero/hello/greet/internal/types"
)

func CreateUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), ctx)
		err := l.CreateUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}