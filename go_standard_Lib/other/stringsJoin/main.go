package main

import (
	"strings"
	"fmt"
)

func main() {
	str := "helloworld"
	str2 := convert(str,3)
	fmt.Println(str2)
}


func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	//设周期为t
	t := 2 * numRows - 2
	rows := make([]string,numRows)

	for k ,v := range s {
		x := k % t
		rows[min(x,t-x)] += string(v)
	}
	return strings.Join(rows,"")
}

func min(a,b int) int {
	if a < b {
		return  a
	}
	return b
}


