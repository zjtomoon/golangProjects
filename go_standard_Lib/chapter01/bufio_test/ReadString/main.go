package main

import (
	"bufio"
	"strings"
	"fmt"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line, _ := reader.ReadString('\n')
	fmt.Printf("the line:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadString('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}

//调用了ReadBytes方法，并将结果的[]byte转为string类型