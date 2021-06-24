package main

import "fmt"

func main() {
	//指定大小的显式初始化
	a := [3]int{1,2,3}

	//通过"..."由后面的元素个数推断数组大小
	b := [...]int{1,2,3}

	//指定大小，并通过索引值初始化，未显式初始化的元素被置为"零值"
	c := [3]int{1:1,2:3}

	//指定大小但不显式初始化，数组元素全被置为"零值"
	var d [3]int
	fmt.Printf("len=%d,value=%v\n",len(a),a) //len=3,value=[1 2 3]
	fmt.Printf("len=%d,value=%v\n",len(b),b) //len=3,value=[1 2 3]
	fmt.Printf("len=%d,value=%v\n",len(c),c) //len=3,value=[0 1 3]
	fmt.Printf("len=%d,value=%v\n",len(d),d) //len=3,value=[0 0 0]
}
