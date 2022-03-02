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

func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	// 获取制作指向的元素
	v = v.Elem()
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("kuteng")
	}
}

func main() {
	u := User{1, "51mh.com", 20}
	SetValue(&u)
	fmt.Println(u)
}
