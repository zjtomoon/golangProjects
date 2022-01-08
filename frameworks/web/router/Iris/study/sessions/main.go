package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func secret(ctx iris.Context) {
	if auth, _ := sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}

	ctx.WriteString("The cake is a lie")
}

func login(ctx iris.Context) {
	session := sess.Start(ctx)

	session.Set("authenticated", true)
}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)

	session.Set("authenticated", false)
	session.Delete("authenticated")
	session.Destroy()
}

func main() {
	app := iris.New()
	app.Get("/secret", secret)
	app.Get("/login", login)
	app.Get("/logout", logout)

	app.Run(iris.Addr(":3000"))
}
