package main

var n int

func foo() (int,error) {
	return 1,nil
}

//访问全局变量n
func g()  {
	println(n)
}

func main() {
	//此时main函数作用域里面没有n
	//所以创建新的局部变量n
	n,_ := foo()

	//访问的全局变量n
	g() //0

	//访问的是main函数作用域下的n
	println(n)
}
