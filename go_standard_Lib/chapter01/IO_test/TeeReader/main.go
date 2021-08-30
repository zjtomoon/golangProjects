package main

import (
	"io"
	"os"
	"strings"
)

/*

TeeReader 返回一个 Reader，它将从 r 中读到的数据写入 w 中。所有经由它处理的
从 r 的读取都匹配于对应的对 w 的写入。它没有内部缓存，即写入必须在读取完成前
完成。任何在写入时遇到的错误都将作为读取错误返回。
也就是说，我们通过 Reader 读取内容后，会自动写入到 Writer 中去

*/

func main() {
	reader := io.TeeReader(strings.NewReader("Go语言中文网"), os.Stdout)
	reader.Read(make([]byte, 20))
}
