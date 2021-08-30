package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, _ := strconv.ParseInt("128", 10, 8)
	fmt.Printf("%d\n", n) //127
}

//如果字符串表示的整数超过了 bitSize 参数能够表示
//的范围，则会返回 ErrRange，同时会返回 bitSize 能够表示的最大或最小值
