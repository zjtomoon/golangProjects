package main

import (
	"io"
	"fmt"
)

/*
	声明新接口类型的特点
	（1）接口的命名一般以“er”结尾
	（2）接口定义的内部方法声明不需要func引导
	（3）在接口定义中，只有方法声明没有方法实现
 */

 type Reader interface {
 	Read(p []byte) (n int,err error)
 }

 type Writer interface {
 	Write(p []byte) (n int,err error)
 }

 //如下3种声明是等价的，最终的展开模式都是第3种格式
/* type ReadWriter interface {
 	Reader
 	Writer
 }*/
/* type ReadWriter interface {
 	Reader
 	Write(p []byte) (n int,err error)
 }*/

 type ReadWriter interface {
 	Read(p []byte) (n int,err error)
 	Write(p []byte) (n int,err error)
 }

 //没有初始化的接口变量，其默认值是nil
 //=========================

type Printer interface {
	Print()
}

type S struct {

}

func (s S) Print()  {
	println("print")
}

func main() {
	var i io.Reader
	fmt.Printf("%T\n",i) //nil

	//=====================
	fmt.Println("==================")
	var j Printer
	//j.Print()
	// panic: runtime error: invalid memory address or nil pointer dereference
	//没有初始化的接口调用其方法会产生panic

	j = S{}
	j.Print() //print

}
