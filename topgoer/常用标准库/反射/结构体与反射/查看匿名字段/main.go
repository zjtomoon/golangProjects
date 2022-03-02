package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// 匿名字段
type Boy struct {
	User
	Addr string
}

func main() {
	m := Boy{User{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}
