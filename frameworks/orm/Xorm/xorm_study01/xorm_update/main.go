package main

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Level   int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	engine.ID(1).Update(&User{Name: "ldj"})
	engine.ID(1).Cols("name", "age").Update(&User{Name: "dj"})
	engine.Table(&User{}).ID(1).Update(map[string]interface{}{"age": 18})
}

/*
	由于使用map[string]interface{}类型的参数，xorm无法推断表名，必须使用Table()方法指定。
	第一个Update()方法只会更新name字段，其他空字段不更新。第二个Update()方法会更新name和age两个字段，age被更新为 0。
*/
