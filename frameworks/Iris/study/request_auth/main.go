package main

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

//Iris 通过它的 jwt 中间件，提供请求权限验证

//jwt 中间件有3个方法来验证 token：
//
//第一个方法是 Serve 方法，这是一个 iris.Handler
//第二个方法是 CheckJWT(iris.Context) bool
//第三个方法是 Get(iris.Context) *jwt.Token，这是一个用于取回已验证的token。

func getTokenHandler(ctx iris.Context) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"foo": "bar",
	})

	tokenString, _ := token.SignedString([]byte("My secret"))

	ctx.HTML(`Token: ` + tokenString + `<br/><br/>
    <a href="/secured?token=` + tokenString + `">/secured?token=` + tokenString + `</a>`)
}

func myAuthenticatedHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)

	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")

	foobar := user.Claims.(jwt.MapClaims)
	for key, value := range foobar {
		ctx.Writef("%s = %s", key, value)
	}
}

func main() {
	app := iris.New()
	j := jwt.New(jwt.Config{
		Extractor: jwt.FromParameter("token"),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("My Secret"), nil
		},
		SigningMethod: jwt.SigningMethodES256,
	})

	app.Get("/", getTokenHandler)
	app.Get("/secured", j.Serve, myAuthenticatedHandler)

	//URL查询参数
	app.Get("/welcome", func(ctx iris.Context) {
		firstname := ctx.URLParamDefault("firstname", "Guest")
		lastname := ctx.URLParam("lastname")

		ctx.Writef("Hello %s %s", firstname, lastname)
	})

	app.Run(iris.Addr(":3000"))
}
