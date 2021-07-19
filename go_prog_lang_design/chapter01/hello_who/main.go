package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {

	who := "World !"

	if len(os.Args) > 1 { /*os.Args[0]是"hello"或"hello.exe" 即可执行文件*/
		who = strings.Join(os.Args[1:]," ")
	}
	fmt.Println("Hello",who)

}