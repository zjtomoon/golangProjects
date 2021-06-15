package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./notes.text", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var i io.Reader = f
	switch v := i.(type) {
	//多个类型，f满足其中任何一个就算匹配
	case *os.File, io.ReadWriter:
		if v == i {
			fmt.Println(true)
		}
	default:
		return
	}
}
