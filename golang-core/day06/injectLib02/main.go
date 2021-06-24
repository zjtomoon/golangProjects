package main
//go get github.com/codegangsta/inject
import (
	"github.com/codegangsta/inject"
	"fmt"
)

type S1 interface {

}

type S2 interface {

}

type Staff struct {
	Name     string `inject`
	Company  S1     `inject`
	Level    S2     `inject`
	Age      int    `inject`
}

func main() {
	//创建被注入实例
	s := Staff{}

	//控制实例的创建
	inj := inject.New()

	//初始化注入值
	inj.Map("tom")
	inj.MapTo("tencent",(*S1)(nil))
	inj.MapTo("T4",(*S2)(nil))
	inj.Map(23)

	//实现对struct注入
	inj.Apply(&s)

	//打印结果
	fmt.Printf("s = %v\n",s)
}
/*
	inject包的处理流程
	(1) 通过inject.New()创建注入引擎，注入引擎被引擎，返回的是Inject接口类型变量
	(2) 调用TypeMapper接口(Injector内嵌TypeMapper)的方法注入struct的字段值或函数的实参值
	(3) 调用Invoker方法执行被注入的函数，或者调用Applicator接口方法获得被注入后的结构实例
 */
