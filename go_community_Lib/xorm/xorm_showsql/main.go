package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	engine.ShowSQL(true)

	user := User{}
	engine.ID(1).Omit("created", "updated").Get(user)
	fmt.Printf("user:%v\n", user)
}
