package main

import (
	"fmt"
	"time"
)

func main() {

	var m = [...]int{1,2,3}

	for i,v := range m {
		//使用临时变量保存当前值

		i := i
		v := v
		go func() {
			fmt.Println(i,v)
		}()
	}

	time.Sleep(3 * time.Second)
}
