package main

import (
	"database/sql"
	"fmt"
	"time"
)

func main() {
	db,_ := sql.Open("mysql","root:@tcp(localhost:3306)/test?charset=utf8")
	fmt.Println("please exec show processlist")
	time.Sleep(10 * time.Second)
	fmt.Println("please exec show processlist again")
	db.Ping()
	time.Sleep(10 * time.Second)
}
