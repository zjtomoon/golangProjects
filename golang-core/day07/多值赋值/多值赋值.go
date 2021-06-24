package main

import "fmt"
/*
	(1) 对左侧操作数中的表达式、索引值进行计算和确定，首先确定左侧的操作数的地址；
		然后对右侧的赋值表达式进行计算，如果发现右侧的表达式计算引用了左侧的变量，
		则创建临时变量进行值拷贝，最后完成计算。
	(2) 从左到右的顺序依次赋值。
 */
func main() {
	x := []int{1,2,3}
	i := 0
	i,x[i] = 1,2 //set i = 1, x[0] = 2
	fmt.Println(i,x)

	x = []int{1,2,3}
	i = 0
	x[i],i = 2,1 //set x[0] = 2 , i = 1
	fmt.Println(i,x)

	x = []int{1,2,3}
	i = 0
	x[i], i = 2 , x[i] //set tmp=x[0],x[0]=2,i = tmp ==> i =1
	fmt.Println(i,x)
	x[0],x[0] = 1,2 //set x[0]=1,then x[0]=2

	fmt.Println(x[0])
}
