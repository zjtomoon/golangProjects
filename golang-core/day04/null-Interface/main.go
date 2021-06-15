package main

import "fmt"

//空接口
type Inter interface {
	Ping()
	Pang()
}

type St struct {
}

func (St) Ping() {
	println("pang")
}

func (*St) Pang() {
	println("pang")
}

func main() {
	var st *St = nil
	var it Inter = st

	fmt.Printf("%p\n", st) //pang
	fmt.Printf("%p\n", it) //0x0

	if it != nil {
		it.Pang() //0x0
		//下面的语句会导致panic
		//方法转换为函数调用，第一个参数是St类型，由于*St是nil,无法获取指针所指的对象值，所以导致panic
		//it.Ping() //panic: value method main.St.Ping called using nil *St pointer
	}
}
