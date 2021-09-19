package main

import "fmt"

func main(){
	arr := [...]int64{1,2,3,4,5,6}
	s1 := arr[0:4]
	fmt.Println(sum1(s1))
}

func sum1(array []int64) int64 {
	var res int64 = 0
	for i:= 0 ; i < len(array) ; i++ {
		res += array[i]
	}
	return res
}