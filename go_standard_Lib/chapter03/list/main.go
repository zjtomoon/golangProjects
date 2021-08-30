package main

import (
	"container/list"
	"fmt"
)

func main() {
	list1 := list.New()
	list1.PushBack(1)
	list1.PushBack(2)

	fmt.Printf("len: %v\n", list1.Len())
	fmt.Printf("first: %#v\n", list1.Front())
	fmt.Printf("second: %#v\n", list1.Front().Next())
}
