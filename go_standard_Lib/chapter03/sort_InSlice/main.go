package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4}

	//对[]int切片排序是更常使用sort.Ints()，而不是直接使用IntSlice类型：
	sort.Ints(s) //[1 2 3 4 5 6]

	//sort.Sort(sort.Reverse(sort.IntSlice(s)))  //[6 5 4 3 2 1]

	//fmt.Println(s)
	//SearchInts()的使用条件为：切片a已经升序排序
	fmt.Println(sort.SearchInts(s, 3)) //2
}
