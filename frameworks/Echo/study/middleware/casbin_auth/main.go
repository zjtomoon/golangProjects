package main

import (
	"github.com/casbin/casbin"
	"github.com/labstack/echo"
	//casbin_mw "github.com/labstack/echo-contrib/casbin"
)

func main() {
	e := echo.New()
	//e.Use(casbin_mw.Middleware(casbin.NewEnforcer("casbin_auth_model.conf", "casbin_auth_policy.csv")))
	ce := casbin.NewEnforcer("casbin_auth_model.conf", "")
	ce.AddRoleForUser("alice", "admin")
	e.Use()
	//ce.AddPolicy(...)
	//e.Use(casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
	//	Enforcer:ce,
	//}))
}
