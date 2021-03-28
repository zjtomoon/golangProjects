package main

func fa(base int) (func(int) int, func(int) int) {
	println(&base, base)
	add := func(i int) int {
		base += i
		println(&base, base)
		return base
	}

	sub := func(i int) int {
		base -= i
		println(&base, base)
		return base
	}

	return add, sub
}

func main() {
	//f g 闭包引用的base是同一个,是fa函数调用传递过来的实参值
	f, g := fa(0)
	//s k 闭包引用的base是同一个,是fa函数调用传递过来的实参值
	s, k := fa(0)
	//f g和s k 引用不同的闭包变量,这是由于fa每次调用都要重新分配形参
	println(f(1), g(2))
	println(s(1), k(2))

}
