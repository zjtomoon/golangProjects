package main

import "fmt"

//类型断言

func main() {
	var x interface{}
	x = "hello"
	v,ok := x.(string)
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("类型断言失败!")
	}
}
func justifyType(x interface{})  {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string ,value is %v\n",v)
	case int:
		fmt.Printf("x is an int, value is %v\n",v)
	case bool:
		fmt.Printf("x is a bool,value is %v\n",v)
	default:
		fmt.Printf("unsupported type")
	}
}