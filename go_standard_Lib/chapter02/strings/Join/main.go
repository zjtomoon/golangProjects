package main

import (
	"bytes"
	"fmt"
)

func Join(str []string, sep string) string {
	//特殊情况应该做处理
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str[0]
	}
	buffer := bytes.NewBufferString(str[0])
	for _, s := range str[1:] {
		buffer.WriteString(sep)
		buffer.WriteString(s)
	}
	return buffer.String()
}

func main() {
	fmt.Println(Join([]string{"name=xxx", "age=xx"}, "&"))
}

//name=xxx&age=xx

//标准库实现
/*func Join(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}
	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}*/
