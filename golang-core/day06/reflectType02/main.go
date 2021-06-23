package main

import "reflect"

type INT int

type A struct {
	a int
}

type B struct {
	b string
}

type Ita interface {
	String() string
}

func (b B) String() string {
	return b.b
}

func main() {
	var a INT = 12
	var b int = 14

	//实参是具体类型，reflect.TypeOf返回是其静态类型
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	//INT和int是两个类型，两者不相等
	if ta == tb {
		println("ta == tb")
	} else {
		println("ta != tb")
	}

	println(ta.Name())
	println(tb.Name())

	//底层基础类型
	println(ta.Kind().String())
	println(tb.Kind().String())

	s1 := A{1}
	s2 := B{"tata"}

	//实参是具体类型，reflect.TypeOf返回的是静态类型
	println(reflect.TypeOf(s1).Name())
	println(reflect.TypeOf(s2).Name())

	//Type的Kind()方法返回的是基础类型，类型A和B的底层基础类型都是struct
	println(reflect.TypeOf(s1).Kind().String())
	println(reflect.TypeOf(s2).Kind().String())

	ita := new(Ita)
	var itb Ita = s2

	//实参是未绑定具体变量的接口类型，reflect.TypeOf返回的是接口类型本身
	//也就是接口的静态类型
	println(reflect.TypeOf(ita).Elem().Name())
	println(reflect.TypeOf(itb).Kind().String())

	println(reflect.TypeOf(itb).Name())
	println(reflect.TypeOf(itb).Kind().String())
	
}