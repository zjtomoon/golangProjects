package main

import (
	"flag"
	"fmt"
	"time"
)

//定义命令行参数对应的变量，这四个变量都是指针类型

var cliName = flag.String("name", "John", "Input your name")
var cliAge = flag.Int("age", 22, "Input your Age")
var cliGender = flag.String("gender", "male", "Input your Gender")
var cliPeriod = flag.Duration("period", 1*time.Second, "sleep period")

//定义一个值类型的命令行参数变量，在Init()函数中对其初始化
//因此，命令行参数对应变量的定义和初始化是可以分开的
var cliFlag int

func Init() {
	flag.IntVar(&cliFlag, "flagname", 1234, "Just fot demo")
}

func main() {
	//初始化变量cliFlag
	Init()
	//把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	//flag.Args() //函数返回没有被解析的命令行参数
	//flag.NArg() //函数返回没有被解析的命令行参数的个数

	fmt.Printf("args = %s,num = %d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d] = %s\n", i, flag.Arg(i))
	}

	//输出命令行参数

	fmt.Println("name = ", *cliName)
	fmt.Println("age = ", *cliAge)
	fmt.Println("gender = ", *cliGender)
	fmt.Printf("period = %v", *cliPeriod)
}

//执行 go run main.go -name=jack
//无flag.Parse()

/*args = [],num = 0
name =  jack
age =  22
gender =  male
period = 1s
*/

//有flag.Parse()
/*
args = [],num = 0
name =  John
age =  22
gender =  male
period = 1s
*/
