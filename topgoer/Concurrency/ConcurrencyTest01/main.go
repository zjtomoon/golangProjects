package main

import (
	"fmt"
	"time"
)


func hello() {
	fmt.Println("Hello,Goroutinue !")
}

func main() {
	go hello()
	fmt.Println("main goroutinue done")
	time.Sleep(time.Second)

}