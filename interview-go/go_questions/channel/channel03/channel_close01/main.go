package main

import "fmt"

func IsClosed(ch <- chan int) bool {
	select {
	case <- ch:
		return true
	default:
	}
	return false
}

func main() {
	c := make(chan int)
	fmt.Println(IsClosed(c)) //false
	close(c)
	fmt.Println(IsClosed(c)) //true
}
