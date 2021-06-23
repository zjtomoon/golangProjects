package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"1111"b:"3333"`
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")

	//获取tag数据
	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")

	/*
	可以像Json一样，取tag里的数据，多个tag之间无逗号，tag不需要引号
	*/
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	fmt.Println("type_number:",rt.Name())
	fmt.Println("type_NumField:",rt.PkgPath())
	fmt.Println("type_String:",rt.String())
	fmt.Println("type.Kind.String:",rt.Kind().String())
	fmt.Println("type.String()=",rt.String())

	//获取结构类型的字段名称
	for i := 0; i < rt.NumField() ; i++ {
		fmt.Printf("type.Field[%d].Name := %v \n",i,rt.Field(i).Name)
	}
	sc := make([]int,10)
	sc = append(sc,1,2,3)
	sct := reflect.TypeOf(sc)

	//获取slice元素的Type
	scet := sct.Elem()

	fmt.Println("slice element type.Kind()=",scet.Kind())
	fmt.Printf("slice element type.Kind()=%d",scet.Kind())
	fmt.Println("slice element type.String()=",scet.String())

	fmt.Println("slice element type.Name()=",scet.Name())
	fmt.Println("slice type.NumMethod()=",scet.NumMethod())
	fmt.Println("slice type.PkgPath()=",scet.PkgPath())
	fmt.Println("slice type.PkgPath()=",sct.PkgPath())
}
