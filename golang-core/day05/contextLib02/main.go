package main

import "context"

type otherContext struct {
	context.Context
}

func main() {
	//使用context.Backgrond()构建一个WithCancel类型的上下文
	ctxa,cancel := context.WithCancel(context.Background())
	//work模拟运行并检测前端的退出通知
	
}

func work()  {
	
}

func workWithValue()  {
	
}