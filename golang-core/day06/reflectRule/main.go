package main

import (
	"reflect"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{1,"andes",20}

	va := reflect.ValueOf(u)
	vb := reflect.ValueOf(&u)
	//fmt.Println(va,vb)

	//值类型是可修改的
	fmt.Println(va.CanSet(),va.FieldByName("Name").CanSet()) //false false

	//指针类型是可以修改的
	fmt.Println(vb.CanSet(),vb.Elem().FieldByName("Name").CanSet()) //false true

	fmt.Printf("%v\n",vb)
	name := "shine"
	vc := reflect.ValueOf(name)

	//通过Set函数修改变量的值
	vb.Elem().FieldByName("Name").Set(vc)
	fmt.Printf("%v\n",vb)

}
/*
 反射三定律
(1) 反射可以从接口值得到反射对象
(2) 反射可以从反射对象获得接口值
(3) 若要修改一个反射对象，则其值必须可以修改
 */