package main

import (
	"io"
	"log"
	"os"
)

//类型查询

func main() {
	f, err := os.OpenFile("./notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var i io.Reader = f
	switch v := i.(type) {
	//i的绑定的实例是*osFile类型，实现了io.ReadWriter接口，所以case匹配成功
	case io.ReadWriter:
		//v是io.ReadWriter接口类型，所以可以调用write方法
		v.Write([]byte("io.ReadWriter\n"))

		//由于上一个case已经匹配，就算这个case也匹配，也不会走到这里
	case *os.File:
		v.Write([]byte("os.File\n"))
		v.Sync()
	default:
		return
	}
}
