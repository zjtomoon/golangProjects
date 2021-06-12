package main

import "fmt"

type Person struct {
	name string
	age int
}

func main(){
	//使用struct字面量声明的是未命名类型
	a := struct {
		name string
		age int
	}{"andes",18}
	fmt.Printf("%T\n",a)
	fmt.Printf("%v\n",a)

	b := Person{"tom",21}
	fmt.Printf("%T\n",b)
	fmt.Printf("%v\n",b)
}
