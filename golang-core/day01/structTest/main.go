package main

import "fmt"

func main(){
	//struct类型变量的初始化
	type Person struct {
		Name string
		Age int
	}
	type Student struct {
		*Person
		Number int
	}
	//推荐这种初始化方式，一旦struct增加字段，则整个初始化语句会报错
	//a := Person{"Tom",21}
	p := &Person{
		Name: "tata",
		Age: 12,
	}
	s := Student {
		Person :p,
		Number: 110,
	}
	fmt.Println(s.Person)
}
