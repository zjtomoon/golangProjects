package main

import "fmt"
//空接口
func main()  {
	var x interface{}
	s := "pprof.cn"
	x = s
	fmt.Printf("type:%T,value:%v\n",x,x)
	i := 100
	x = i
	fmt.Printf("type:%T,value:%v\n",x,x)
	b := true
	x = b
	fmt.Printf("type:%T,value:%v\n",x,x)
}
