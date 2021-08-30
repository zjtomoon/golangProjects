package main

import (
	"fmt"
	"time"
)

func main() {
	t1, _ := time.Parse("2006-01-02 15:04:05", "2021-08-30 18:10:00")
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
	fmt.Println(t2)
	fmt.Println(time.Now().Sub(t1).Hours())
}
