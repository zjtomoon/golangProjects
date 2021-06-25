package main

func fa(a int) func(i int) int {
	return func(i int) int {
		println(&a,a)
		a = a + i
		return a
	}
}

func main() {
	// f是一个闭包，包括对函数fa形式参数a的"同名引用"
	f := fa(1)
	println(f(1))
	println(f(2))
}
