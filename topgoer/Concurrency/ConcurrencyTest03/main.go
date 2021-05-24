package main

import (
	"fmt"
	"time"
)


func main() {
	go func ()  {
		i := 0
		for {
			i++
			fmt.Printf("new goroutinue: i = %d\n",i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutinue: i = %d\n",i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}