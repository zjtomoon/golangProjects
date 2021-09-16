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
	engine, _ := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")

	slcUsers := make([]User, 1) //切片
	engine.Where("age > ? and age < ?", 12, 30).Find(&slcUsers)
	fmt.Println("users whose age between [12,30]:", slcUsers)

	mapUsers := make(map[int64]User)
	engine.Where("length(name) = ?", 3).Find(&mapUsers)
	fmt.Println("users whose has name of length 3:", mapUsers)
}
