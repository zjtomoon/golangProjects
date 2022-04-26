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

	engine.Where("age > ? and age < ?", 12, 30).Iterate(&User{}, func(i int, bean interface{}) error {
		fmt.Printf("user%d:%v\n", i, bean.(*User))
		return nil
	})
}

/*
	Find()一样，Iterate()也是找到满足条件的所有记录，只不过传入了一个回调去处理每条记录
 */