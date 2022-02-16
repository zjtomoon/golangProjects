package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("https://studygolang.com \nIt is the home of gophers"))
	line, _ := reader.ReadBytes('\n')
	fmt.Printf("the line : %s\n", line)
	//这里可以换上任意的bufio的Read/Write操作
	n, _ := reader.ReadBytes('\n')
	fmt.Printf("the line : %s \n", n)
	fmt.Println(string(n))
}

//ReadBytes返回的值不会是指向Reader中的buffer
