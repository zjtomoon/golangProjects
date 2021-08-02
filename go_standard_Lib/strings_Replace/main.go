package main

import (
	"strings"
	"fmt"
)

func main() {
	s := "hello world "
	s2 := replaceSpace(s)
	fmt.Println(s2)
}

func replaceSpace(s string) string  {
	return strings.Replace(s," ","%20",-1)
}