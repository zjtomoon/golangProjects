package main

import (
	"bufio"
	"strings"
	"fmt"
	"time"
)

func main() {
	reader := bufio.NewReaderSize(strings.NewReader("http://studygolang.com.\t It is the home of gophers"),14)
	go Peek(reader)
	go reader.ReadBytes('\t')
	time.Sleep(1e8)
}

func Peek(reader *bufio.Reader)  {
	line,_ := reader.Peek(14)
	fmt.Printf("%s\n",line)
	time.Sleep(1)
	fmt.Printf("%s\n",line)
}
//Reader的Peek方法如果返回的[]byte⻓度小于n，这时返回的err为非nil，用于
// 解释为啥会小于n。如果n大于reader的buffer⻓度，err会是ErrBufferFull。