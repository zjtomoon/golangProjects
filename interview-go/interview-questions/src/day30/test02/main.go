package main

import "fmt"

func main() {
	var a = [5]int{1,2,3,4,5}
	var r [5]int

	for i,v := range a {
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
	range 表达式是副本参与循环，就是说例子中参与循环的是 a 的副本，而不是真正的 a。
	因此无论 a 被如何修改，其副本 b 依旧保持原值，并且参与循环的是 b，
	因此 v 从 b 中取出的仍旧是 a 的原值，而非修改后的值。
 */