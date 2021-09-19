package main

import (
	"fmt"
)
func sum(arr ...int) (sum int){
	for _, v:= range arr {
		sum += v
	}
	return 
}

func suma(arr ...int) (sum int){
	for v := range arr {
		sum += v
	}
	return 
}
func sumb(arr []int) (sum int){
	for v := range arr {
		sum += v
	}
	return
}
func main(){
	slice := []int{1,2,3,4}
	//array := [...]int{1,2,3,4}
	//数组不可以作为实参传递
	res :=sum(slice...)
	fmt.Println(res)

	fmt.Println("============")
	//suma和sumb的类型并不一样
	fmt.Printf("%T\n",suma)
	fmt.Printf("%T\n",sumb)
}