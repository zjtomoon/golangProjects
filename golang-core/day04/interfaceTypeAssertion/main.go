package main

import "fmt"

//接口类型断言

type Inter interface {
	Ping()
	Pang()
}
type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping()  {
	println("ping")
}
func (*St) Pang()  {
	println("pang")
}

func main() {
	st := &St{"andes"}
	var i interface{} = st

	//判断i 绑定的实例是否实现了接口类型Inter
	o := i.(Inter)
	o.Ping()
	o.Pang()

	//如下语句会引发panic,因为i没有实现接口Anter
	//p := i.(Anter)
	//p.String()
	// panic: interface conversion: *main.St is not main.Anter: missing method String

	//判断i绑定的实例是否就是具体类型St
	s := i.(*St)
	fmt.Printf("%s\n",s.Name)
}
