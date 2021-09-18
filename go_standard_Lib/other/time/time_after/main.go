package main

import (
	"time"
	"fmt"
)

func main() {
	tchan := time.After(3 * time.Second)
	fmt.Printf("tchan type = %T \n ",tchan)
	fmt.Println("mark 1")

	fmt.Println("tchan = ",<- tchan) //等待参数duration时间后，向返回的chan里面写入当前时间。
	fmt.Println("mark 2")
}
