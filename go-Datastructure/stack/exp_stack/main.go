package main

import (
	"errors"
	"fmt"
)

//使用栈的思想模拟计算器

//使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop 	int // 表示我们栈最大可以存放数个数
	Top		int //表示栈顶，因为栈顶固定，因此我们直接使用Top
	arr		[20]int //数组模拟栈
}

//入栈
func (this *Stack) Push(val int) (err error) {
	
	//先判断栈是否满了
	if this.Top == this.MaxTop - 1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

//出栈
func (this *Stack) Pop() (val int, err error) {
	//判断栈是否空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0,errors.New("stack empty")
	}

	//先取值，再this.Top--
	val = this.arr[this.Top]
	this.Top--
	return val,nil
}

//遍历栈，注意需要从栈顶开始遍历
func (this *Stack) List()  {
	//先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
}

func main()  {
	
}