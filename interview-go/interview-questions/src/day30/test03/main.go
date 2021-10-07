package main

import "fmt"

func main() {
	var a = [5]int{1,2,3,4,5}
	var r [5]int

	for i,v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ",r)
	fmt.Println("a = ",a)
}

/*
	修复代码中，使用 *[5]int 作为 range 表达式，其副本依旧是一个指向原数组 a 的指针，
	因此后续所有循环中均是 &a 指向的原数组亲自参与的，因此 v 能从 &a 指向的原数组中取出 a 修改后的值。
 */