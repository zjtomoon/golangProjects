package main

import (
	"io"
	"os"
	"fmt"
	"strings"
)

func main() {
	io.Copy(os.Stdout,os.Stdin)
	fmt.Println("Got EOF --- bye")
	io.CopyN(os.Stdout,strings.NewReader("Go语言中文网"),8)
}
