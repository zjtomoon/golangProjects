package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn failed,error :", err)
		return
	}
	defer c.Close()

	//请求/响应服务可以实现持续处理新请求，客户端可以发送多个命令到服务器而无需等待响应，最后在一次读取多个响应。
	//
	//使用Send()，Flush()，Receive()方法支持管道化操作
	//
	//Send向连接的输出缓冲中写入命令。
	//
	//Flush将连接的输出缓冲清空并写入服务器端。
	//
	//Recevie按照FIFO顺序依次读取服务器的响应

	c.Send("Set", "name1", "sss1")
	c.Send("Set", "name2", "sss2")

	c.Flush()

	v, err := c.Receive()
	fmt.Printf("v: %v,err:%v\n", v, err)

	v, err = c.Receive()
	fmt.Printf("v: %v,err:%v\n", v, err)

	v, err = c.Receive() // 夯住，一直等待

	fmt.Printf("v: %v,err:%v\n", v, err)
}
