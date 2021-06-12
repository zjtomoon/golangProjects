//类型强制转换
package main

import "fmt"

type Map map[string]string

func (m Map) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

type iMap Map

//只要底层类型是slice、map等支持range的类型字面量，新类型仍然可以使用range迭代

func (m iMap) Print() {
	for _, key := range m {
		fmt.Println(key)
	}
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"
	//mp与ma有相同的底层类型map[string]string，并且mp是未命名类型
	var ma Map = mp
	//im与ma虽然有相同的底层类型，但是二者中没有一个是字面量类型，不能直接赋值，可以强制进行类型转换
	//var im iMap = ma
	var im iMap = (iMap)(ma)
	ma.Print()
	im.Print()
}

/*
	(1)数值类型和string类型之间的相互转换可能造成值部分丢失；其他的转换仅仅是类型的转换，不会造成值的改变
	string和数字之间的转换可使用标准库strconv
	(2)GO语言没有语言机制支持指针和integer之间的直接转换，可以使用标准库中的unsafe包进行处理
*/
