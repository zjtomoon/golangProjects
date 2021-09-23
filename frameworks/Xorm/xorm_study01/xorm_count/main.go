package main

import (
	"fmt"
	"time"
	"xorm.io/xorm"
	_ "github.com/go-sql-driver/mysql"
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
	engine,_ := xorm.NewEngine("mysql","root:123@/test?charset=utf8")

	num,_ := engine.Where("age >= ?",50).Count(&User{})
	fmt.Printf("there are %d users whose age >= 50",num)
}
