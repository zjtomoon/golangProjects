package main

import (
	"database/sql"
	"time"
)

func main() {
	db,_ := sql.Open("mysql","root:@tcp(localhost:3306)/test?charset=utf8")
	//去掉注释，可以看看相应的空闲连接是不是变化了
	db.SetMaxIdleConns(3)
	for i := 0 ; i < 10 ; i++ {
		go func() {
			db.Ping()
		}()
	}
	time.Sleep(20 * time.Second)
}
