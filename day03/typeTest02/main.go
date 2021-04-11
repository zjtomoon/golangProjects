//类型相同和类型赋值
package main

import "fmt"

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

//只要底层类型是slice、map等支持range的类型字面值，新类型的字面量，新类型仍然可以使用range迭代
func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type slice []int

func (s slice) Print() {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"

	//mp和ma有相同的底层类型map[string]string,并且map是未命名类型
	//所以mp可以直接赋值给ma
	var ma Map = mp
	//im与ma虽然有相同的底层类型map[string]string，但是它们中没有一个是未命名类型
	//不能赋值，如下语句不能通过编译
	//var im iMap = ma
	ma.Print()
	//im.Print()

	var i interface {
		Print()
	} = ma

	i.Print()

	s1 := []int{1, 2, 3}
	var s2 slice
	s2 = s1
	s2.Print()
}

//p78
