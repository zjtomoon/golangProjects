package main

import (
	"fmt"
	"time"
)

//future模式
/*
	future模式的基本工作原理
	(1) 使用chan 作为函数参数
	(2) 启动goroutinue调用函数
	(3) 通过chan传入参数
	(4) 做其他可以并行处理的事情
	(5) 通过chan异步获取结果
*/

//一个查询结构体
//这里的sql和result是一个简单的抽象，具体的应用可能是更复杂的数据类型
type query struct {
	//参数Channel
	sql chan string
	//结果Channel
	result chan string
}

//执行Query
func execQuery(q query)  {
	//启动协程
	go func() {
		//获取输入
		sql := <-q.sql
		//访问数据库

		//输出结果通道
		q.result <- "result from " + sql
	}()
}


func main() {
	//初始化Query
	q := query{make(chan string,1),make(chan string,1)}
	//执行Query，注意执行的时候无须准备参数
	go execQuery(q)
	//发送参数
	q.sql <- "select * from table"
	//做其他事情，通过sleep描述
	time.Sleep(1 * time.Second)

	//获取结果
	fmt.Println(<-q.result)
}
