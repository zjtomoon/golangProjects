package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com.\nIt is the home of gophers"))
	line, _ := reader.ReadSlice('\n')
	fmt.Printf("the line : %s \n", line)
	//这里可以换上任意的bufio的Read/Write操作
	n, _ := reader.ReadSlice('\n')
	fmt.Printf("the line : %s\n", line)
	fmt.Println(string(n))
}
