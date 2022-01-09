package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//在事务开启的时候，我们就去redis中setnx一条数据，这条数据的键要和你当前操作的数据有关，
// 这样就只会锁定一条数据，而不影响其他数据的业务，例如：做订单审核的时候，将订单号+业务简写作为键。
//判断上面插入操作的返回值，如果返回1，就继续执行，如果返回0，直接return.
//在事务结束之后，将redis中的这条数据删除。直接使用del（String key）就可以了。

//管道操作可以理解为并发操作，并通过Send()，Flush()，Receive()三个方法实现。
// 客户端可以使用send()方法一次性向服务器发送一个或多个命令，命令发送完毕时，
// 使用flush()方法将缓冲区的命令输入一次性发送到服务器，客户端再使用Receive()方法依次按照先进先出的顺序读取所有命令操作结果。

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn failed,error :", err)
		return
	}
	defer c.Close()

	c.Send("HSET", "student", "name", "wd", "age", "22")
	c.Send("HSET", "student", "Score", "100")
	c.Send("HGET", "student", "age")
	c.Flush()

	res, err := c.Receive()
	fmt.Printf("Receive res: %v\n", res)

	res2, err := c.Receive()
	fmt.Printf("Receive res2:%v\n", res2)

	res3, err := c.Receive()
	fmt.Printf("Receive res3:%s\n", res3)
}
