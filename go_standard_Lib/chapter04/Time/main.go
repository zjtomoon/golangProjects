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

	t3,_ := time.ParseInLocation("2006-01-02 15:04:05",time.Now().Format("2006-01-02 15:04:05"),time.Local)
	//整点 （向下取整）
	fmt.Println(t3.Truncate(1 * time.Hour))
	//整点 （最接近）
	fmt.Println(t3.Round(1 * time.Hour))

	//整分 （向下取整）
	fmt.Println(t3.Truncate(1 * time.Minute))
	//整分 （最接近）
	fmt.Println(t3.Round(1 * time.Minute))

	t4,_ := time.ParseInLocation("2006-01-02 15:04:05",t3.Format("2006-01-02 15:04:05"),time.Local)
	fmt.Println(t4)
}
