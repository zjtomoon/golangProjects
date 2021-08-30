package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	/*	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) //Println will add back the final '\n'
	}
	if err := scanner.Err();err != nil {
		fmt.Fprintln(os.Stderr,"reading standard input:",err)
	}*/

	file, err := os.Create("./scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("http://studygolang.com.\nIt is the home of gopher.\nIf you are studying golang,welcome you!")
	//将文件offset设置到文件头
	file.Seek(0, os.SEEK_SET)
	scanner := bufio.NewReader(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
