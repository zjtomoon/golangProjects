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

	//事务操作

	//MULTI, EXEC,DISCARD和WATCH是构成Redis事务的基础，当然我们使用go语言对redis进行事务操作的时候本质也是使用这些命令。
	//
	//MULTI：开启事务
	//
	//EXEC：执行事务
	//
	//DISCARD：取消事务
	//
	//WATCH：监视事务中的键变化，一旦有改变则取消事务。

	c.Send("MULTI")
	c.Send("INCR", "foo")
	c.Send("INCR", "bar")
	r, err := c.Do("EXEC")

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println(r)
}
