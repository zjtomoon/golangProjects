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
	Level   int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine,_ := xorm.NewEngine("mysql","root:123@/test?charset=utf8")

	querySql := "select * from user limit 1"
	results,_ := engine.Query(querySql)
	for _,record := range results {
		for key,val := range record {
			fmt.Println(key,string(val))
		}
	}

	updateSql := "update `user` set name=? where id=?"
	res,_ := engine.Exec(updateSql,"ldj",1)
	fmt.Println(res.RowsAffected())
}
