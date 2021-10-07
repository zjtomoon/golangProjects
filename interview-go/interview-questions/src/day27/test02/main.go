package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo":Math{2,3},
}

func main() {
	/*m["foo"].x = 4
	fmt.Println(m["foo"].x)*/
	/*
		对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的。
	 */
	 tmp := m["foo"]
	 tmp.x = 4
	 m["foo"] = tmp
	 fmt.Println(m["foo"].x)

}
