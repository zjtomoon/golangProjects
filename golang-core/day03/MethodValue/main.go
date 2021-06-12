//方法值
/*
	方法值是一个带有闭包的函数变量，其底层实现原理和带有闭包的匿名函数类似，
	接收值被隐式地绑定到方法值的闭包环境中，后续调用不需要再显式地传递接收者。
*/
package main

import (
	"fmt"
)

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(i int) {
	t.a = i
}
func (t *T) Print() {
	fmt.Printf("%p,%v,%d \n", t, t, t.a)
}

func main() {
	var t = &T{}

	//method value
	f := t.Set

	f(2)

	t.Print()

	f(3)
	t.Print()

}
