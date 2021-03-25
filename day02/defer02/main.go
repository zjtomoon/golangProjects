package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("defer")
	}()
	println("func body")
	os.Exit(1)
}

/*
	defer的好处是可以在一定程度上避免资源泄露。。特别是有很多return语句
*/
