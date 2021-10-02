package main

import "fmt"

func main() {
	/*list := new([]int)
	list = append(list,1)*/
	//无法通过编译

	list := make([]int,0)
	list = append(list,1)
	fmt.Println(list)
}
