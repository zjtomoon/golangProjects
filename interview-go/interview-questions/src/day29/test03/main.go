package main

import (
	"fmt"
	"time"
)

func main() {

	var m = [...]int{1, 2, 3}

	for i, v := range m {
		//使用函数传递
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}

	time.Sleep(3 * time.Second)
}
