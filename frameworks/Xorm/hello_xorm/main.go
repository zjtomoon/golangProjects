package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	var err error
	engine,err = xorm.NewEngine("mysql","root:123/test?charset=utf-8")
	if err != nil {
		panic(err)
	}
}
