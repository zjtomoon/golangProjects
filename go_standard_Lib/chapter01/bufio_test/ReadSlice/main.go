package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	/*	reader := bufio.NewReader(strings.NewReader("http://studygolang.com.\nIt is the home of gophers"))
		line, _ := reader.ReadSlice('\n')
		fmt.Printf("the line : %s \n", line)
		//这里可以换上任意的bufio的Read/Write操作
		n, _ := reader.ReadSlice('\n')
		fmt.Printf("the line : %s\n", line)
		fmt.Println(string(n))*/

	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com"), 16)
	line, err := reader.ReadSlice('\n')
	fmt.Printf("line: %s \t error: %s\n", line, err)
	line, err = reader.ReadSlice('\n')
	fmt.Printf("line: %s \t error : %s \n", line, err)
}

//ReadSlice返回的[]byte是指向Reader中的buffer,而不是copy一份返回
