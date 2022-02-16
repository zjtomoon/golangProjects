package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _, _ := reader.ReadLine()
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _, _ := reader.ReadLine()
	fmt.Printf("the line:%s\n", n)
	fmt.Println(string(n))
}
