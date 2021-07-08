package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup

func add() {
	for i := 0 ; i < 5000 ; i++ {
		x += 1
	}
	wg.Add(-1)
	//wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
