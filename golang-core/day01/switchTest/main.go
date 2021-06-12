package main

import "fmt"

func main(){
	switch i := "y"; i {
	case "y","Y"://多个case值使用逗号分隔
		fmt.Println("yes")
		fallthrough
		//fallthrough会跳过接下来的case条件表达式直接执行下一个case语句
	case "n" , "N":
		fmt.Println("no")
	}
}
